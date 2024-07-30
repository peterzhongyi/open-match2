// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package store

import "regexp"

// An enum for the type of operations that the replication queue can process.
const (
	Ticket = iota
	Activate
	Deactivate
	Assign
)

// Every change to the state of the tickets in om-core is modelled as a StateUpdate.
type StateUpdate struct {
	Cmd   int    // The operation this update contains
	Key   string // The key to update
	Value string // The value to associate with this key (if applicable)
}

// Results of changes to the state of the cache. State replication batches
// updates as much as possible, and every update generates a StateResponse that
// can be sent back to the underlying caller (since it may get batched with
// unrelated updates generated by RPC calls that happen concurrently). This is
// primarily used as the mechanism to send back a response from a CreateTicket
// call - this is how that call receives the ticket ID from the state storage
// layer. It may also used in internal implementations to track which updates
// have been applied to the local replicated ticket cache.
//
// If err is nil, result contains the replication id assigned to the update by the
// state storage implementation (in redis, this will be a stream event ID, for
// example).
//
// If err is not nil, result contains the key of the StateUpdate that failed,
// used by the calling function to track what requests caused errors.
type StateResponse struct {
	Result string
	Err    error
}

// The core gRPC server instantiates a replicatedTicketCache on startup, and specifies
// how it wants to replicate om-core state by instantiating a StateReplicator
// that conforms to this interface.
type StateReplicator interface {
	GetUpdates() []*StateUpdate
	SendUpdates([]*StateUpdate) []*StateResponse
	GetReplIdValidator() *regexp.Regexp
}