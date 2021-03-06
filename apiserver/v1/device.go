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
	"net/http"
	"time"

	"github.com/cloustone/sentel/apiserver/db"
	"github.com/golang/glog"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// Device internal definition
type registerDeviceRequest struct {
	requestBase
	ProductKey string `bson:"productKey"`
	DeviceName string `bson:"productName"`
}

type registerDeviceResponse struct {
	DeviceId     string    `bson:"deviceId"`
	DeviceName   string    `bson:"deviceName"`
	DeviceSecret string    `bson:"deviceSecret"`
	DeviceStatus string    `bson:"deviceStatus"`
	ProductKey   string    `bson:"productKey"`
	TimeCreated  time.Time `bson:"timeCreated"`
}

// RegisterDevice register a new device in IoT hub
// curl -d "ProductKey=7&DeviceName=2" "http://localhost:4145/api/v1/devices/3?api-version=v1"
func registerDevice(ctx echo.Context) error {
	// Get product
	glog.Infof("RegisterDevice ctx:[%s] :%s, %s ", ctx, ctx.ParamNames(), ctx.ParamValues())
	req := new(registerDeviceRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response{Success: false, Message: err.Error()})
	}
	config := ctx.(*apiContext).config
	// Connect with registry
	r, err := db.NewRegistry(config)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response{Success: false, Message: err.Error()})
	}
	defer r.Release()

	//          
	// Insert device into registry, the created product
	// will be modified to retrieve specific information sucha as
	// product.id and creation time
	dp := db.Device{
		Id:           uuid.NewV4().String(),
		Name:         req.DeviceName,
		ProductKey:   req.ProductKey,
		TimeCreated:  time.Now(),
		TimeModified: time.Now(),
	}
	err = r.RegisterDevice(&dp)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response{Success: false, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, &response{Success: true,
		Result: &registerDeviceResponse{
			DeviceId:     dp.Id,
			DeviceName:   dp.Name,
			ProductKey:   dp.ProductKey,
			DeviceSecret: dp.DeviceSecret,
			TimeCreated:  dp.TimeCreated,
		}})
}

// Retrieve a device from the identify registry of an IoT hub
// curl -XDELETE "http://localhost:4145/api/v1/devices/3?api-version=v1"
func getDevice(ctx echo.Context) error {
	glog.Infof("getDevice ctx:[%s] :%s, %s ", ctx, ctx.ParamNames(), ctx.QueryParams())
	// Get product
	req := new(registerDeviceRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response{Success: false, Message: err.Error()})

	}
	// Connect with registry
	config := ctx.(*apiContext).config
	registry, err := db.NewRegistry(config)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response{Success: false, Message: err.Error()})
	}
	defer registry.Release()

	// Get device into registry, the created product
	dev, err := registry.GetDevice(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusOK,
			&response{RequestId: uuid.NewV4().String(), Success: false, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK,
		&response{
			Success: true,
			Result: &registerDeviceResponse{
				DeviceId:     dev.Id,
				DeviceName:   dev.Name,
				ProductKey:   dev.ProductKey,
				DeviceSecret: dev.DeviceSecret,
				TimeCreated:  dev.TimeCreated,
				DeviceStatus: dev.DeviceStatus,
			}})
}

// Delete the identify of a device from the identity registry of an IoT Hub
// curl -XDELETE "http://localhost:4145/api/v1/devices/3?api-version=v1"
func deleteDevice(ctx echo.Context) error {
	// Get product
	req := new(registerDeviceRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response{Success: false, Message: err.Error()})
	}
	// Connect with registry
	config := ctx.(*apiContext).config
	registry, err := db.NewRegistry(config)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response{Success: false, Message: err.Error()})
	}
	defer registry.Release()

	rcp := &response{
		RequestId: uuid.NewV4().String(),
		Success:   true,
	}
	// Get device into registry, the created product
	if err := registry.DeleteDevice(ctx.Param("id")); err != nil {
		return ctx.JSON(http.StatusOK, &response{Success: false, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, rcp)
}

type updateDeviceRequest struct {
	requestBase
	DeviceId     string `bson:"deviceId"`
	DeviceName   string `bson:"productName"`
	ProductKey   string `bson:"productKey"`
	DeviceSecret string `bson:"deviceSecrt"`
	DeviceStatus string `bson:"deviceStatus"`
}
type updateDeviceResponse struct {
	DeviceId     string    `bson:"deviceId"`
	DeviceName   string    `bson:"deviceName"`
	DeviceSecret string    `bson:deviceSecret"`
	DeviceStatus string    `bson:deviceStatus"`
	ProductKey   string    `bson:"productKey"`
	TimeCreated  time.Time `bson:"timeCreated"`
	TimeModified time.Time `bson:"timeModified"`
}

// updateDevice update the identity of a device in the identity registry of an IoT Hub
// curl -XPUT -d "ProductKey=7&DeviceName=2" "http://localhost:4145/api/v1/devices/3?api-version=v1"
func updateDevice(ctx echo.Context) error {
	// Get product
	req := new(updateDeviceRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response{Message: err.Error()})
	}
	// Connect with registry
	r, err := db.NewRegistry(ctx.(*apiContext).config)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response{Message: err.Error()})
	}
	defer r.Release()

	// Insert device into registry, the created product
	// will be modified to retrieve specific information sucha as
	// product.id and creation time
	dp := db.Device{
		Name:         req.DeviceName,
		ProductId:    ctx.Param("id"),
		ProductKey:   req.ProductKey,
		DeviceStatus: req.DeviceStatus,
		DeviceSecret: req.DeviceSecret,
		TimeModified: time.Now(),
	}
	if err = r.UpdateDevice(&dp); err != nil {
		return ctx.JSON(http.StatusOK,
			&response{RequestId: uuid.NewV4().String(), Success: false, Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK,
		&response{Success: true,
			Result: &updateDeviceResponse{
				DeviceId:     dp.Id,
				DeviceName:   dp.Name,
				ProductKey:   dp.ProductKey,
				DeviceSecret: dp.DeviceSecret,
				TimeCreated:  dp.TimeCreated,
				TimeModified: dp.TimeModified,
			}})
}

// Delete all the pending commands for this devices from the IoT hub
func purgeCommandQueue(ctx echo.Context) error {
	return nil
}

// Query an IoT hub to retrieve information regarding device twis using a SQL-like language
func queryDevices(ctx echo.Context) error {
	return nil
}

// Get the identifies of multiple devices from The IoT hub
func getMultipleDevices(ctx echo.Context) error {
	return nil
}
