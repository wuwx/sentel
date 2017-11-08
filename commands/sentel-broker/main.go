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

package main

import (
	"flag"

	"github.com/cloustone/sentel/broker"
	"github.com/cloustone/sentel/broker/api"
	"github.com/cloustone/sentel/broker/base"
	"github.com/cloustone/sentel/broker/metric"
	"github.com/cloustone/sentel/broker/mqtt"
	"github.com/cloustone/sentel/core"

	"github.com/golang/glog"
)

var (
	configFileFullPath = flag.String("c", "stentel-broker.conf", "config file")
)

func main() {
	var mgr *base.ServiceManager
	var config core.Config
	var err error

	flag.Parse()
	glog.Info("Starting mqtt broker...")

	// Check all registered service
	if err := base.CheckAllRegisteredServices(); err != nil {
		glog.Fatal(err)
		return
	}
	// Get configuration
	if config, err = core.NewWithConfigFile(*configFileFullPath); err != nil {
		glog.Fatal(err)
		flag.PrintDefaults()
		return
	}
	// Create service manager according to the configuration
	if mgr, err = base.NewServiceManager(config); err != nil {
		glog.Fatal("Failed to launch ServiceManager")
		return
	}
	glog.Error(mgr.Run())
}

func init() {
	for group, values := range broker.DefaultConfigs {
		core.RegisterConfig(group, values)
	}
	base.RegisterService("mqtt:tcp", mqtt.Configs, &mqtt.MqttFactory{})
	base.RegisterService("mqtt:ssl", mqtt.Configs, &mqtt.MqttFactory{})
	base.RegisterService("mqtt:ws", mqtt.Configs, &mqtt.MqttFactory{})
	base.RegisterService("api", api.Configs, &api.ApiServiceFactory{})
	base.RegisterService("metric", metric.Configs, &metric.MetricServiceFactory{})
}
