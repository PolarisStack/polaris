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

package config

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mockcache "github.com/polarismesh/polaris/cache/mock"
	"github.com/polarismesh/polaris/common/eventhub"
	mockstore "github.com/polarismesh/polaris/store/mock"
	"github.com/stretchr/testify/assert"
)

func Test_Initialize(t *testing.T) {
	eventhub.InitEventHub()
	ctrl := gomock.NewController(t)
	mockStore := mockstore.NewMockStore(ctrl)
	cacheMgr := mockcache.NewMockCacheManager(ctrl)

	t.Cleanup(func() {
		ctrl.Finish()
		eventhub.Shutdown()
		originServer.watchCenter.Close()
		originServer.initialized = false
	})

	cacheMgr.EXPECT().OpenResourceCache(gomock.Any()).Return(nil).AnyTimes()
	cacheMgr.EXPECT().ConfigFile().Return(nil).AnyTimes()
	cacheMgr.EXPECT().Gray().Return(nil).AnyTimes()
	cacheMgr.EXPECT().ConfigGroup().Return(nil).AnyTimes()

	err := Initialize(context.Background(), Config{
		Open: true,
	}, mockStore, cacheMgr, nil)
	assert.NoError(t, err)
	assert.NotNil(t, originServer)
}
