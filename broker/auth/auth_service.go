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

package auth

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/cloustone/sentel/core"

	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ServiceName = "auth"
)

// AuthServiceFactory
type AuthServiceFactory struct{}

// New create coap service factory
func (p *AuthServiceFactory) New(c core.Config, quit chan os.Signal) (core.Service, error) {
	// check mongo db configuration
	hosts, _ := core.GetServiceEndpoint(c, "broker", "mongo")
	timeout := c.MustInt("broker", "connect_timeout")
	session, err := mgo.DialWithTimeout(hosts, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// Connect with redis if cache policy is redis
	var rclient *redis.Client = nil
	addr, _ := core.GetServiceEndpoint(c, "broker", "redis")
	password := c.MustString("broker", "redis_password")
	db := c.MustInt("auth", "redis_db")

	rclient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if _, err := rclient.Ping().Result(); err != nil {
		return nil, err
	}

	return &AuthService{
		ServiceBase: core.ServiceBase{
			Config:    c,
			Quit:      quit,
			WaitGroup: sync.WaitGroup{},
		},
		rclient: rclient,
	}, nil
}

// Authentication Service
type AuthService struct {
	core.ServiceBase
	rclient *redis.Client
}

// Name
func (p *AuthService) Name() string {
	return ServiceName
}

// Start
func (p *AuthService) Start() error {
	return nil
}

// Stop
func (p *AuthService) Stop() {
	signal.Notify(p.Quit, syscall.SIGINT, syscall.SIGQUIT)
	p.WaitGroup.Wait()
	close(p.Quit)
}

// CheckAcl check client's access control right
func (p *AuthService) authorize(clientid string, username string, topic string, access int, opt *Options) error {
	return nil
}

// authenticate check user's name and password
func (p *AuthService) authenticate(opt *Options) error {
	if key, err := p.getDeviceSecretKey(opt); err == nil {
		opt.DeviceSecret = key
		return sign(opt)
	}
	return fmt.Errorf("auth: Failed to get device secret key for '%s'", opt.DeviceName)
}

// Device
type device struct {
	Id           string    `bson:"Id"`
	Name         string    `bson:"Name"`
	ProductId    string    `bson:"ProductId"`
	ProductKey   string    `bson:"productKey"`
	DeviceStatus string    `bson:"deviceStatus"`
	DeviceSecret string    `bson:"deviceSecret"`
	TimeCreated  time.Time `bson:"timeCreated"`
	TimeModified time.Time `bson:"TimeModified"`
}

// getDeviceSecretKey retrieve device secret key from cache or mongo
func (p *AuthService) getDeviceSecretKey(opt *Options) (string, error) {
	// Read from cache at first
	key := opt.ProductKey + "/" + opt.DeviceName
	if val, err := p.rclient.Get(key).Result(); err == nil {
		return val, nil
	}

	// Read from database if not found in cache
	hosts, _ := core.GetServiceEndpoint(p.Config, "broker", "mongo")
	timeout := p.Config.MustInt("broker", "connect_timeout")
	session, err := mgo.DialWithTimeout(hosts, time.Duration(timeout)*time.Second)
	if err != nil {
		return "", err
	}
	defer session.Close()
	c := session.DB("registry").C("devices")

	dev := device{}
	if err := c.Find(bson.M{"ProductKey": opt.ProductKey, "DeviceName": opt.DeviceName}).One(&dev); err != nil {
		return "", err
	}
	// Write back to redis
	p.rclient.Set(key, dev.DeviceSecret, 0)
	return dev.DeviceSecret, nil
}
