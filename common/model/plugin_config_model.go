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

package model

import "errors"

// DiscoverEventConfig 服务实例事件插件配置
type DiscoverEventConfig struct {
	QueueSize     int                         `json:"queueSize"`
	LoggerConfigs []DiscoverEventLoggerConfig `json:"logger"`
}

type DiscoverEventLoggerConfig struct {
	Name   string                 `json:"name"`
	Option map[string]interface{} `json:"option"`
}

// Validate 检查配置是否正确配置
func (c *DiscoverEventConfig) Validate() error {
	if c.QueueSize <= 0 {
		return errors.New("QueueSize is <= 0")
	}
	if len(c.LoggerConfigs) == 0 {
		return errors.New("LoggerConfig is empty")
	}
	return nil
}
