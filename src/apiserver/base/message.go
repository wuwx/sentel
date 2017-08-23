//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use this file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.

package base

import (
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
)

func SyncProduceMessage(c echo.Context, topic string, value sarama.Encoder) error {
	ctx := *c.(*ApiContext)

	//	sarama.Logger = c.Logger()

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(c.Request().RemoteAddr),
		Value: value,
	}

	producer, err := sarama.NewSyncProducer(strings.Split(ctx.Config.Kafka, ","), config)
	if err != nil {
		c.Logger().Error("Failed to produce message:%s", err)
		return err
	}
	defer producer.Close()

	if _, _, err := producer.SendMessage(msg); err != nil {
		c.Logger().Error("Failed to send producer message:%s", err)
	}
	return err
}

func AsyncProduceMessage(c echo.Context, topic string, value sarama.Encoder) error {
	ctx := *c.(*ApiContext)

	//	sarama.Logger = c.Logger()

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(c.Request().RemoteAddr),
		Value: value,
	}

	producer, err := sarama.NewAsyncProducer(strings.Split(ctx.Config.Kafka, ","), config)
	if err != nil {
		c.Logger().Error("Failed to produce message:%s", err)
		return err
	}
	defer producer.Close()

	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					c.Logger().Error(err)
				}
			case <-success:
			}
		}
	}(producer)

	producer.Input() <- msg
	return err
}
