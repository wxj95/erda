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

package kafka

import (
	"fmt"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	writer "github.com/erda-project/erda-infra/pkg/parallel-writer"
	"github.com/erda-project/erda-infra/providers/kafka"
	"github.com/erda-project/erda/modules/oap/collector/common"
	"github.com/erda-project/erda/modules/oap/collector/core/model/odata"
	"github.com/erda-project/erda/modules/oap/collector/plugins"
)

var providerName = plugins.WithPrefixExporter("kafka")

type config struct {
	MetadataKeyOfTopic string               `file:"metadata_key_of_topic"`
	Producer           kafka.ProducerConfig `file:"producer"`
	// capability of old data format
	Compatibility bool `file:"compatibility" default:"true"`
}

// +provider
type provider struct {
	Cfg *config
	Log logs.Logger

	Kafka  kafka.Interface `autowired:"kafka"`
	writer writer.Writer
}

func (p *provider) ComponentConfig() interface{} {
	return p.Cfg
}

func (p *provider) Connect() error {
	return nil
}

func (p *provider) Export(ods []odata.ObservableData) error {
	for _, item := range ods {
		data, err := p.serialize(item, p.Cfg.Compatibility)
		if err != nil {
			p.Log.Errorf("serialize err: %s", err)
			continue
		}

		if p.Cfg.MetadataKeyOfTopic != "" {
			tmp, ok := item.Metadata().Get(p.Cfg.MetadataKeyOfTopic)
			if !ok {
				p.Log.Errorf("unable to find topic with key %s", p.Cfg.MetadataKeyOfTopic)
				continue
			}

			if err := p.writer.Write(&kafka.Message{
				Topic: &tmp,
				Data:  data,
			}); err != nil {
				p.Log.Errorf("write data to %s err: %s", tmp, err)
			}
		} else {
			if err := p.writer.Write(data); err != nil {
				p.Log.Errorf("write data to %s err: %s", p.Cfg.Producer.Topic, err)
			}
		}
	}
	return nil
}

func (p *provider) serialize(od odata.ObservableData, compatibility bool) ([]byte, error) {
	if od.SourceType() == odata.RawType {
		return od.Source().([]byte), nil
	}

	data, err := common.JSONSerializeSingle(od, compatibility)
	if err != nil {
		return nil, fmt.Errorf("JSONSerializeSingle item err: %s", err)
	}
	return data, nil
}

// Run this is optional
func (p *provider) Init(ctx servicehub.Context) error {
	producer, err := p.Kafka.NewProducer(&p.Cfg.Producer)
	if err != nil {
		return err
	}
	p.writer = producer

	return nil
}

func init() {
	servicehub.Register(providerName, &servicehub.Spec{
		Services: []string{
			providerName,
		},
		Description: "here is description of erda.oap.collector.exporter.kafka",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
