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

package event

import "github.com/cloustone/sentel/broker/broker"

type EventHandler func(e *Event, ctx interface{})

// Publish publish event to event service
func Notify(e *Event) {
	e.BrokerId = broker.GetId()
	service := broker.GetService(ServiceName).(*EventService)
	service.notify(e)
}

// Subscribe subcribe event from event service
func Subscribe(event uint32, handler EventHandler, ctx interface{}) {
	service := broker.GetService(ServiceName).(*EventService)
	service.subscribe(event, handler, ctx)
}
