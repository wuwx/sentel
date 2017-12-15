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

package v1

import (
	"github.com/cloustone/sentel/apiserver/base"
	"github.com/cloustone/sentel/apiserver/middleware"
	"github.com/cloustone/sentel/libs/sentel"
)

func NewApi(c sentel.Config) *base.ApiManager {
	m := base.NewApiManager("v1", c)

	// Tenant Api
	m.RegisterApi(base.POST, "/tenants/:id", addTenant)
	m.RegisterApi(base.DELETE, "/tenants/:id", deleteTenant, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/tenants/:id", getTenant, middleware.DefaultKeyAuth)

	// Product Api
	m.RegisterApi(base.POST, "/products/:id", registerProduct, middleware.DefaultKeyAuth)
	m.RegisterApi(base.DELETE, "/products/:id", deleteProduct, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/products/:id", getProduct, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/products/:id/devices", getProductDevices, middleware.DefaultKeyAuth)

	// Device Api
	m.RegisterApi(base.POST, "/devices/:id", registerDevice, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/devices/:id", getDevice, middleware.DefaultKeyAuth)
	m.RegisterApi(base.DELETE, "/devices/:id", deleteDevice, middleware.DefaultKeyAuth)
	m.RegisterApi(base.PUT, "/devices/:id", updateDevice, middleware.DefaultKeyAuth)
	m.RegisterApi(base.DELETE, "/devices/:id/commands", purgeCommandQueue, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/devices/", getMultipleDevices, middleware.DefaultKeyAuth)
	m.RegisterApi(base.POST, "/devices/query", queryDevices, middleware.DefaultKeyAuth)

	// Statics Api
	m.RegisterApi(base.GET, "/statistics/devices", getRegistryStatistics, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/statistics/service", getServiceStatistics, middleware.DefaultKeyAuth)

	// Device Twin Api
	m.RegisterApi(base.GET, "/twins/:id", getDeviceTwin, middleware.DefaultKeyAuth)
	m.RegisterApi(base.POST, "/twins/:id/methods", invokeDeviceMethod, middleware.DefaultKeyAuth)
	m.RegisterApi(base.PATCH, "/twins/:id", updateDeviceTwin, middleware.DefaultKeyAuth)

	// Http Runtime Api
	m.RegisterApi(base.POST, "/devices/:id/messages/deviceBound/:etag/abandon",
		abandonDeviceBoundNotification, middleware.DefaultKeyAuth)
	m.RegisterApi(base.DELETE, "/devices/:id/messages/devicesBound/:etag",
		completeDeviceBoundNotification, middleware.DefaultKeyAuth)
	m.RegisterApi(base.POST, "/devices/:ideviceId/files",
		createFileUploadSasUri, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/devices/:id/message/deviceBound",
		receiveDeviceBoundNotification, middleware.DefaultKeyAuth)
	m.RegisterApi(base.POST, "/devices/:deviceId/files/notifications",
		updateFileUploadStatus, middleware.DefaultKeyAuth)
	m.RegisterApi(base.POST, "/devices/:id/messages/event", sendDeviceEvent, middleware.DefaultKeyAuth)

	// Job Api
	m.RegisterApi(base.POST, "/jobs/:jobid/cancel", cancelJob, middleware.DefaultKeyAuth)
	m.RegisterApi(base.PUT, "/jobs/:jobid", createJob, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/jobs/:jobid", getJob, middleware.DefaultKeyAuth)
	m.RegisterApi(base.GET, "/jobs/query", queryJobs, middleware.DefaultKeyAuth)

	return m
}
