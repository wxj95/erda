// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clickhouse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	log "github.com/erda-project/erda/modules/core/monitor/log"
)

func Test_fillLogInfo(t *testing.T) {
	p := &provider{}

	logData := &log.Log{
		Tags: map[string]string{
			"dice_org_name": "erda",
			"msp_env_id":    "env_id",
		},
		ID:        "123456789012345",
		Timestamp: time.Now().UnixNano(),
	}

	p.fillLogInfo(logData)

	assert.Equal(t, "env_id", logData.TenantId)
	assert.Equal(t, "erda", logData.OrgName)
}
