//
// Copyright 2019 Insolar Technologies GmbH
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
//

package transport

import (
	"context"
	"fmt"
	"time"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/heavy/replica"
	"github.com/insolar/insolar/network"
	"github.com/insolar/insolar/network/hostnetwork/host"
	"github.com/insolar/insolar/network/hostnetwork/packet"
	"github.com/insolar/insolar/network/hostnetwork/packet/types"
	"github.com/pkg/errors"
)

func NewHostTransport(hostNetwork network.HostNetwork) replica.Transport {
	handlers := make(map[string]replica.Handle)
	registerHandlers(hostNetwork, handlers)
	return &hostTransport{
		handlers: handlers,
		net:      hostNetwork,
	}
}

type hostTransport struct {
	net      network.HostNetwork
	handlers map[string]replica.Handle
}

func (t *hostTransport) Send(ctx context.Context, receiver, method string, data []byte) ([]byte, error) {
	receiverHost, err := t.hostByAddress(receiver)
	if err != nil || receiverHost == nil {
		return nil, errors.Wrapf(err, "failed to create host by receiver address")
	}
	req := &packet.RPCRequest{
		Method: method,
		Data:   data,
	}
	future, err := t.net.SendRequestToHost(context.Background(), types.Replication, req, receiverHost)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send request to host")
	}
	packet, err := future.WaitResponse(10 * time.Second)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get result from host")
	}
	if packet.GetResponse() == nil || packet.GetResponse().GetRPC() == nil {
		return nil, errors.Errorf("error getting RPC response from node %s: got invalid response protobuf message: %s", receiver, packet)
	}
	resp := packet.GetResponse().GetRPC()
	if resp.Result == nil {
		return nil, errors.Errorf("RPC call returned error: %s", resp.Error)
	}
	return resp.Result, nil
}

func (t *hostTransport) Register(method string, handle replica.Handle) {
	t.handlers[method] = handle
}

func (t *hostTransport) Me() string {
	return t.net.PublicAddress()
}

func registerHandlers(net network.HostNetwork, handlers map[string]replica.Handle) {
	net.RegisterRequestHandler(types.Replication, func(ctx context.Context, req network.Packet) (network.Packet, error) {
		if req.GetRequest() == nil || req.GetRequest().GetRPC() == nil {
			inslogger.FromContext(ctx).Warnf("process RPC: got invalid request protobuf message: %s", req)
		}

		method := req.GetRequest().GetRPC().Method
		data := req.GetRequest().GetRPC().Data
		if _, ok := handlers[method]; !ok {
			return net.BuildResponse(ctx, req, &packet.RPCResponse{
				Error: fmt.Sprintf("handle function: %v not defined", method),
			}), nil
		}
		result, err := handlers[method](data)
		reply, err := insolar.Serialize(replica.GenericReply{
			Data:  result,
			Error: err,
		})
		if err != nil {
			return net.BuildResponse(ctx, req, &packet.RPCResponse{Error: err.Error()}), nil
		}
		return net.BuildResponse(ctx, req, &packet.RPCResponse{Result: reply}), nil
	})
}

func (t *hostTransport) hostByAddress(receiver string) (*host.Host, error) {
	host, err := host.NewHost(receiver)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create new host by address")
	}
	return host, nil
}
