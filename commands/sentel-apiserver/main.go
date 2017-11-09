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

	"github.com/cloustone/sentel/core"

	"github.com/cloustone/sentel/apiserver/v1"

	"github.com/cloustone/sentel/apiserver/base"
	"github.com/cloustone/sentel/apiserver/db"

	"github.com/golang/glog"
)

var (
	configFileFullPath = flag.String("c", "apiserver.conf", "config file")
)

func main() {
	var config core.Config
	var err error

	flag.Parse()
	glog.Info("Starting api server...")

	// Get configuration
	if config, err = core.NewWithConfigFile(*configFileFullPath); err != nil {
		glog.Fatal(err)
		flag.PrintDefaults()
		return
	}

	// Register Api Manager
	base.RegisterApiManager(v1.NewApi(config))

	// Initialize registry
	if err := db.InitializeRegistry(config); err != nil {
		glog.Error("Registry initialization failed:%v", err)
		return
	}

	// Create api manager using configuration
	apiManager, err := base.CreateApiManager(config)
	if err != nil {
		glog.Error("ApiManager creation failed:%v", err)
		return
	}
	glog.Error(apiManager.Start())
}

func init() {
	for group, values := range defaultConfigs {
		core.RegisterConfig(group, values)
	}
}