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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/polarismesh/polaris/common/model"
)

func TestResolveServiceName(t *testing.T) {
	svcId := "1234"
	mockSvc := &model.Service{
		ID:        svcId,
		Namespace: DefaultNamespace,
		Name:      "testSvc",
	}
	svr := &EurekaServer{svcResolver: func(s string) *model.Service {
		if s == svcId {
			return mockSvc
		}
		return nil
	}}
	event := &model.InstanceEvent{SvcId: svcId}
	svr.resolveServiceName(event)
	assert.Equal(t, mockSvc.Namespace, event.Namespace)
	assert.Equal(t, mockSvc.Name, event.Service)

	event = &model.InstanceEvent{SvcId: "xyz"}
	svr.resolveServiceName(event)
	assert.Equal(t, "", event.Namespace)
	assert.Equal(t, "", event.Service)

}
