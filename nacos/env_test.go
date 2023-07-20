// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nacos

import (
	"testing"

	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/stretchr/testify/assert"
)

// TestEnvFunc test env func
func TestEnvFunc(t *testing.T) {
	cpc := &ConfigParamConfig{
		Category:          "retry",
		ServerServiceName: "svc",
		ClientServiceName: "cli",
	}

	assert.Equal(t, int64(8848), NacosPort())
	assert.Equal(t, "127.0.0.1", NacosAddr())
	assert.Equal(t, "", NacosNameSpaceId())
	assert.Equal(t, vo.ConfigParam{
		Type:    vo.JSON,
		Group:   NACOS_DEFAULT_CONFIG_GROUP,
		Content: defaultContent,
		DataId:  "cli.svc.retry",
	}, NaocsConfigParam(cpc))

	t.Setenv(NACOS_ENV_NAMESPACE_ID, "ns")
	t.Setenv(NACOS_ENV_SERVER_ADDR, "1.1.1.1")
	t.Setenv(NACOS_ENV_PORT, "80")
	t.Setenv(NACOS_ENV_CONFIG_DATA_ID, "{{.ClientServiceName}}")
	t.Setenv(NACOS_ENV_CONFIG_GROUP, "{{.Category}}")

	assert.Equal(t, int64(80), NacosPort())
	assert.Equal(t, "1.1.1.1", NacosAddr())
	assert.Equal(t, "ns", NacosNameSpaceId())
	assert.Equal(t, vo.ConfigParam{
		Type:    vo.JSON,
		Group:   "retry",
		Content: defaultContent,
		DataId:  "cli",
	}, NaocsConfigParam(cpc))
}