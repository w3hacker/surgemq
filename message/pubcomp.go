// Copyright (c) 2014 Dataence, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message

// The PUBCOMP Packet is the response to a PUBREL Packet. It is the fourth and
// final packet of the QoS 2 protocol exchange.
type PubcompMessage struct {
	PubackMessage
}

var _ Message = (*PubcompMessage)(nil)

// NewPubcompMessage creates a new PUBCOMP message.
func NewPubcompMessage() *PubcompMessage {
	msg := &PubcompMessage{}
	msg.SetType(PUBCOMP)

	return msg
}