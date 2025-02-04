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

package server

import (
	"sync"

	"github.com/erda-project/erda/apistructs"
)

var (
	clientMutex sync.Mutex
	clientDatas = map[apistructs.ClusterDialerClientType]apistructs.ClusterDialerClientMap{}
)

func updateClientDetail(clientType apistructs.ClusterDialerClientType, clusterKey string, data apistructs.ClusterDialerClientDetail) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	if clusterKey == "" || clientType == "" {
		return
	}
	if _, ok := clientDatas[clientType]; !ok {
		clientDatas[clientType] = apistructs.ClusterDialerClientMap{}
	}
	clientDatas[clientType][clusterKey] = data
}

func getClientDetail(clientType apistructs.ClusterDialerClientType, clusterKey string) (apistructs.ClusterDialerClientDetail, bool) {
	clientMutex.Lock()
	defer clientMutex.Unlock()
	if data, ok := clientDatas[clientType]; ok {
		if clientData, existed := data[clusterKey]; existed {
			return clientData, true
		}
	}
	return apistructs.ClusterDialerClientDetail{}, false
}
