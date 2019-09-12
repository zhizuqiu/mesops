// MIT License
//
// Copyright (c) [2017-2018] [Demitri Swan]
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package v1

import (
	"context"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/mesos/go-proto/mesos/v1/master"
)

// ReserveResource reserves resources dynamically on a specific mesos_v1_agent.
func (m *Master) ReserveResource(ctx context.Context, call *mesos_v1_master.Call_ReserveResources) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_RESERVE_RESOURCES
	var message proto.Message = &mesos_v1_master.Call{Type: &callType, ReserveResources: call}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}

// UnreserveResource unreserves resources dynamically on a specific mesos_v1_agent.
func (m *Master) UnreserveResource(ctx context.Context, call *mesos_v1_master.Call_UnreserveResources) (err error) {
	var callType mesos_v1_master.Call_Type = mesos_v1_master.Call_UNRESERVE_RESOURCES
	var message proto.Message = &mesos_v1_master.Call{Type: &callType, UnreserveResources: call}
	var httpResponse *http.Response
	httpResponse, err = m.client.makeCall(ctx, message, nil)
	defer httpResponse.Body.Close()
	return
}
