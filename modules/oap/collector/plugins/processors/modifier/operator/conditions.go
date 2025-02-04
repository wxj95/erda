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

package operator

import (
	"regexp"
)

type KeyExist struct {
	cfg ConditionCfg
}

func (k *KeyExist) Match(pairs map[string]interface{}) bool {
	_, ok := pairs[k.cfg.Key]
	return ok
}

func NewKeyExist(cfg ConditionCfg) Condition {
	return &KeyExist{cfg: cfg}
}

type ValueMatch struct {
	cfg     ConditionCfg
	pattern *regexp.Regexp
}

func (v *ValueMatch) Match(pairs map[string]interface{}) bool {
	value, ok := pairs[v.cfg.Key]
	if !ok {
		return false
	}
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	return v.pattern.MatchString(strValue)
}

func NewValueMatch(cfg ConditionCfg) Condition {
	return &ValueMatch{cfg: cfg, pattern: regexp.MustCompile(cfg.Value)}
}

type ValueEmpty struct {
	cfg ConditionCfg
}

func (ve *ValueEmpty) Match(pairs map[string]interface{}) bool {
	value, ok := pairs[ve.cfg.Key]
	if !ok {
		return false
	}
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	return strValue == ""
}

func NewValueEmpty(cfg ConditionCfg) Condition {
	return &ValueEmpty{cfg: cfg}
}

type NoopCondition struct {
	cfg ConditionCfg
}

func (n *NoopCondition) Match(pairs map[string]interface{}) bool {
	return true
}

func NewNoopCondition(cfg ConditionCfg) Condition {
	return &NoopCondition{cfg: cfg}
}
