//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.

package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cloustone/sentel/conductor/executor"
	"github.com/cloustone/sentel/core"

	"github.com/Shopify/sarama"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

// Metaservice manage broker metadata
// Broker's metadata include the following data
// - Global broker cluster data
// - Shadow device
type MetadataService struct {
	core.ServiceBase
	consumer sarama.Consumer
}

const (
	MetadataServiceName = "metadata"
)

// MetadataServiceFactory
type MetadataServiceFactory struct{}

// New create metadata service factory
func (p *MetadataServiceFactory) New(c core.Config, quit chan os.Signal) (core.Service, error) {
	// check mongo db configuration
	hosts, _ := core.GetServiceEndpoint(c, "broker", "mongo")
	timeout := c.MustInt("broker", "connect_timeout")
	session, err := mgo.DialWithTimeout(hosts, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// kafka
	khosts, _ := core.GetServiceEndpoint(c, "broker", "kafka")
	consumer, err := sarama.NewConsumer(strings.Split(khosts, ","), nil)
	if err != nil {
		return nil, fmt.Errorf("Connecting with kafka:%s failed", khosts)
	}

	return &MetadataService{
		ServiceBase: core.ServiceBase{
			Config:    c,
			WaitGroup: sync.WaitGroup{},
			Quit:      quit,
		},
		consumer: consumer,
	}, nil

}

// Name
func (p *MetadataService) Name() string {
	return MetadataServiceName
}

// Start
func (p *MetadataService) Start() error {
	return nil
}

// Stop
func (p *MetadataService) Stop() {
	p.consumer.Close()
	p.WaitGroup.Wait()
}

// subscribeTopc subscribe topics from apiserver
func (p *MetadataService) subscribeTopic(topic string) error {
	partitionList, err := p.consumer.Partitions(topic)
	if err != nil {
		return fmt.Errorf("Failed to get list of partions:%v", err)
		return err
	}

	for partition := range partitionList {
		pc, err := p.consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			glog.Errorf("Failed  to start consumer for partion %d:%s", partition, err)
			continue
		}
		defer pc.AsyncClose()
		p.WaitGroup.Add(1)

		go func(sarama.PartitionConsumer) {
			defer p.WaitGroup.Done()
			for msg := range pc.Messages() {
				p.handleNotifications(string(msg.Topic), msg.Value)
			}
		}(pc)
	}
	return nil
}

type ruleTopic struct {
	RuleName  string `json:"ruleName"`
	RuleId    string `json:"ruleId"`
	ProductId string `json:"productId"`
	Action    string `json:"action"`
}

// handleNotifications handle notification from kafka
func (p *MetadataService) handleNotifications(topic string, value []byte) error {
	rule := ruleTopic{}
	if err := json.Unmarshal(value, &topic); err != nil {
		return err
	}
	r := &executor.Rule{
		RuleName:  rule.RuleName,
		RuleId:    rule.RuleId,
		ProductId: rule.ProductId,
		Action:    rule.Action,
	}
	return executor.HandleRuleNotification(r)
}

// getShadowDeviceStatus return shadow device's status
func (p *MetadataService) getShadowDeviceStatus(clientId string) (*Device, error) {
	return nil, nil
}

// syncShadowDeviceStatus synchronize shadow device's status
func (p *MetadataService) syncShadowDeviceStatus(clientId string, d *Device) error {
	return nil
}