/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package eurekaserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gopkg.in/yaml.v2"

	"github.com/polarismesh/polaris/auth"
	"github.com/polarismesh/polaris/cache"
	cachetypes "github.com/polarismesh/polaris/cache/api"
	commonlog "github.com/polarismesh/polaris/common/log"
	"github.com/polarismesh/polaris/common/utils"
	"github.com/polarismesh/polaris/namespace"
	"github.com/polarismesh/polaris/plugin"
	_ "github.com/polarismesh/polaris/plugin/cmdb/memory"
	_ "github.com/polarismesh/polaris/plugin/discoverevent/local"
	_ "github.com/polarismesh/polaris/plugin/healthchecker/memory"
	_ "github.com/polarismesh/polaris/plugin/healthchecker/redis"
	_ "github.com/polarismesh/polaris/plugin/history/logger"
	_ "github.com/polarismesh/polaris/plugin/password"
	_ "github.com/polarismesh/polaris/plugin/ratelimit/lrurate"
	_ "github.com/polarismesh/polaris/plugin/ratelimit/token"
	_ "github.com/polarismesh/polaris/plugin/statis/logger"
	_ "github.com/polarismesh/polaris/plugin/statis/prometheus"
	"github.com/polarismesh/polaris/service"
	"github.com/polarismesh/polaris/service/batch"
	"github.com/polarismesh/polaris/service/healthcheck"
	"github.com/polarismesh/polaris/store"
	storemock "github.com/polarismesh/polaris/store/mock"
	testdata "github.com/polarismesh/polaris/test/data"
)

type Bootstrap struct {
	Logger map[string]*commonlog.Options
}

type TestConfig struct {
	Bootstrap    Bootstrap          `yaml:"bootstrap"`
	Cache        cache.Config       `yaml:"cache"`
	Namespace    namespace.Config   `yaml:"namespace"`
	Naming       service.Config     `yaml:"naming"`
	HealthChecks healthcheck.Config `yaml:"healthcheck"`
	Store        store.Config       `yaml:"store"`
	Auth         auth.Config        `yaml:"auth"`
	Plugin       plugin.Config      `yaml:"plugin"`
}

type EurekaTestSuit struct {
	cfg                 *TestConfig
	server              service.DiscoverServer
	healthSvr           *healthcheck.Server
	namespaceSvr        namespace.NamespaceOperateServer
	cancelFlag          bool
	updateCacheInterval time.Duration
	cancel              context.CancelFunc
	storage             store.Store
	cacheMgr            *cache.CacheManager
}

type options func(cfg *TestConfig)

// 内部初始化函数
func (d *EurekaTestSuit) initialize(t *testing.T, callback func(t *testing.T, s *storemock.MockStore) error, opts ...options) error {
	if err := d.loadConfig(); err != nil {
		return err
	}

	for i := range opts {
		opts[i](d.cfg)
	}

	_ = commonlog.Configure(d.cfg.Bootstrap.Logger)

	commonlog.GetScopeOrDefaultByName(commonlog.DefaultLoggerName).SetOutputLevel(commonlog.ErrorLevel)
	commonlog.GetScopeOrDefaultByName(commonlog.NamingLoggerName).SetOutputLevel(commonlog.ErrorLevel)
	commonlog.GetScopeOrDefaultByName(commonlog.ConfigLoggerName).SetOutputLevel(commonlog.ErrorLevel)
	commonlog.GetScopeOrDefaultByName(commonlog.StoreLoggerName).SetOutputLevel(commonlog.ErrorLevel)
	commonlog.GetScopeOrDefaultByName(commonlog.AuthLoggerName).SetOutputLevel(commonlog.ErrorLevel)

	plugin.SetPluginConfig(&d.cfg.Plugin)

	ctx, cancel := context.WithCancel(context.Background())
	d.cancel = cancel

	// 初始化存储层
	ctrl := gomock.NewController(t)
	s := storemock.NewMockStore(ctrl)
	d.storage = s

	if err := callback(t, s); err != nil {
		return err
	}

	// 初始化缓存模块
	cacheMgn, err := cache.TestCacheInitialize(ctx, &d.cfg.Cache, d.storage)
	if err != nil {
		return err
	}
	d.cacheMgr = cacheMgn

	// 批量控制器
	namingBatchConfig, err := batch.ParseBatchConfig(d.cfg.Naming.Batch)
	if err != nil {
		return err
	}
	healthBatchConfig, err := batch.ParseBatchConfig(d.cfg.HealthChecks.Batch)
	if err != nil {
		return err
	}

	batchConfig := &batch.Config{
		Register:         namingBatchConfig.Register,
		Deregister:       namingBatchConfig.Register,
		ClientRegister:   namingBatchConfig.ClientRegister,
		ClientDeregister: namingBatchConfig.ClientDeregister,
		Heartbeat:        healthBatchConfig.Heartbeat,
	}

	bc, err := batch.NewBatchCtrlWithConfig(d.storage, cacheMgn, batchConfig)
	if err != nil {
		eurekalog.Errorf("new batch ctrl with config err: %s", err.Error())
		return err
	}
	bc.Start(ctx)

	if len(d.cfg.HealthChecks.LocalHost) == 0 {
		d.cfg.HealthChecks.LocalHost = utils.LocalHost // 补充healthCheck的配置
	}
	healthCheckServer, err := healthcheck.TestInitialize(ctx, &d.cfg.HealthChecks, d.cfg.Cache.Open, bc, d.storage)
	if err != nil {
		return err
	}
	cacheProvider, err := healthCheckServer.CacheProvider()
	if err != nil {
		return err
	}
	healthCheckServer.SetServiceCache(cacheMgn.Service())
	healthCheckServer.SetInstanceCache(cacheMgn.Instance())

	// 为 instance 的 cache 添加 健康检查的 Listener
	cacheMgn.AddListener(cachetypes.CacheInstance, []cachetypes.Listener{cacheProvider})
	cacheMgn.AddListener(cachetypes.CacheClient, []cachetypes.Listener{cacheProvider})

	d.healthSvr = healthCheckServer
	time.Sleep(5 * time.Second)
	return nil
}

// 加载配置
func (d *EurekaTestSuit) loadConfig() error {
	d.cfg = new(TestConfig)

	confFileName := testdata.Path("eureka_apiserver_test.yaml")
	if os.Getenv("STORE_MODE") == "sqldb" {
		fmt.Printf("run store mode : sqldb\n")
		confFileName = testdata.Path("eureka_apiserver_test_sqldb.yaml")
	}
	buf, err := ioutil.ReadFile(confFileName)
	if nil != err {
		return fmt.Errorf("read file %s error", confFileName)
	}

	if err = parseYamlContent(string(buf), d.cfg); err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return err
	}
	return err
}

func parseYamlContent(content string, conf *TestConfig) error {
	if err := yaml.Unmarshal([]byte(replaceEnv(content)), conf); nil != err {
		return fmt.Errorf("parse yaml %s error:%w", content, err)
	}
	return nil
}

// replaceEnv replace holder by env list
func replaceEnv(configContent string) string {
	return os.ExpandEnv(configContent)
}

func (d *EurekaTestSuit) Destroy() {
	d.cancel()
	d.cacheMgr.Close()
	d.storage.Destroy()
	healthcheck.TestDestroy()
}
