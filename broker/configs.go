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

package broker

var defaultConfigs = map[string]map[string]string{
	"broker": {
		"loglevel":        "debug",
		"kafka":           "localhost:9092",
		"mongo":           "localhost:27017",
		"connect_timeout": "5",
	},
	"storage": {
		"repository": "local",
		"loglevel":   "debug",
	},
	"auth": {
		"cafile":              "",
		"capath":              "",
		"certfile":            "",
		"keyfile":             "",
		"require_certificate": "false",
	},
	"mqtt": {
		"protocols": "tcp,ws,tls",
		"tcp":       "localhost:1883",
		"tls":       "localhost:1885",
	},
	"rpc": {
		"listen": "localhost:55001",
	},
	"http": {
		"listen": "localhost:55001",
	},
}
