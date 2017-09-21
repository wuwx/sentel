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

package collector

import "encoding/json"

// Notification objects from iothub node

// Subscription
type Subscription struct {
	Ip        string `json:"ip"`
	Name      string `json:"name"`
	CreatedAt string `json:"name"`
	encoded   []byte
	err       error
}

func (p *Subscription) ensureEncoded() {
	if p.encoded == nil && p.err == nil {
		p.encoded, p.err = json.Marshal(p)
	}
}

func (p *Subscription) Length() int {
	p.ensureEncoded()
	return len(p.encoded)
}

func (p *Subscription) Encode() ([]byte, error) {
	p.ensureEncoded()
	return p.encoded, p.err
}