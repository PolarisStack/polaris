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
	"sync"

	"github.com/polarismesh/polaris-server/common/log"
	"go.uber.org/zap"
)

// Event 事件对象，包含类型和事件消息
type Event struct {
	EventType string
	Message   interface{}
}

// Callback 事件回调
type Callback func(event Event) bool

// Center 事件中心
type Center struct {
	watchers map[string][]Callback
	lock     sync.RWMutex
}

// NewEventCenter 新建事件中心
func NewEventCenter() *Center {
	center := &Center{
		watchers: make(map[string][]Callback),
	}

	return center
}

// WatchEvent 监听事件
func (c *Center) WatchEvent(eventType string, cb Callback) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.watchers[eventType]) == 0 {
		c.watchers[eventType] = make([]Callback, 0, 6)
	}

	c.watchers[eventType] = append(c.watchers[eventType], cb)
}

func (c *Center) handleEvent(e Event) {
	defer c.recovery()
	c.lock.RLock()
	defer c.lock.RUnlock()

	cbs, ok := c.watchers[e.EventType]
	if !ok {
		return
	}

	for _, cb := range cbs {
		if !cb(e) {
			log.ConfigScope().Errorf("[Common][Event] cb message error. event = %+v", e)
		}
	}
}

func (c *Center) recovery() {
	if err := recover(); err != nil {
		log.ConfigScope().Error("[Common][Event] handler event error.", zap.Any("error", err))
	}
}
