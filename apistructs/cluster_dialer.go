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

package apistructs

import (
	"encoding/json"
	"fmt"
)

type ClusterDialerHeaderKey string

var (
	ClusterDialerHeaderKeyClusterKey    ClusterDialerHeaderKey = "X-Erda-Cluster-Key"
	ClusterDialerHeaderKeyClientType    ClusterDialerHeaderKey = "X-Erda-Client-Type"
	ClusterDialerHeaderKeyClusterInfo   ClusterDialerHeaderKey = "X-Erda-Cluster-Info"
	ClusterDialerHeaderKeyAuthorization ClusterDialerHeaderKey = "Authorization"
	ClusterDialerHeaderKeyClientDetail  ClusterDialerHeaderKey = "X-Erda-Client-Detail"
)

func (c ClusterDialerHeaderKey) String() string {
	return string(c)
}

type ClusterDialerClientType string

var (
	ClusterDialerClientTypeDefault  ClusterDialerClientType = ""         // cluster
	ClusterDialerClientTypeCluster  ClusterDialerClientType = "cluster"  // cluster
	ClusterDialerClientTypePipeline ClusterDialerClientType = "pipeline" // pipeline
)

func (c ClusterDialerClientType) String() string {
	return string(c)
}

func (c ClusterDialerClientType) MakeClientKey(clusterKey string) string {
	if c == "" {
		return clusterKey
	}
	return fmt.Sprintf("%s-client-type-%s", clusterKey, c)
}

type ClusterDialerClientDetailKey string

var (
	ClusterDialerDataKeyPipelineHost ClusterDialerClientDetailKey = "pipelineHost"
	ClusterDialerDataKeyPipelineAddr ClusterDialerClientDetailKey = "pipelineAddr"
)

type ClusterDialerClientDetail map[ClusterDialerClientDetailKey]string

type ClusterDialerClientMap map[string]ClusterDialerClientDetail

func (detail ClusterDialerClientDetail) Get(key ClusterDialerClientDetailKey) string {
	return detail[key]
}

func (detail ClusterDialerClientDetail) Marshal() ([]byte, error) {
	return json.Marshal(detail)
}

func (m ClusterDialerClientMap) GetClientDetail(clientKey string) ClusterDialerClientDetail {
	return m[clientKey]
}
