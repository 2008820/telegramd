/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc_util

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"github.com/nebulaim/telegramd/mtproto"
	"fmt"
	"context"
	"google.golang.org/grpc/metadata"
	"time"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/nebulaim/telegramd/zproto"
	"github.com/nebulaim/telegramd/grpc_util/service_discovery"
	"github.com/coreos/etcd/clientv3"
	"github.com/nebulaim/telegramd/grpc_util/service_discovery/etcd3"
	"github.com/nebulaim/telegramd/grpc_util/load_balancer"
)

const (
	random = "random"
	round_robin = "round_robin"
	consistent_hash = "consistent_hash"
)

func NewRPCClientByServiceDiscovery(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *grpc.ClientConn, err error) {

	//etcdConfg := clientv3.Config{
	//	Endpoints: []string{"http://127.0.0.1:2379"},
	//}
	//r := etcd3.NewResolver("/nebulaim", "auth_key", etcdConfg)
	//b := load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	//c, err := grpc.Dial("", grpc.WithInsecure(),  grpc.WithBalancer(b), grpc.WithTimeout(time.Second*5))
	//if err != nil {
	//	log.Printf("grpc dial: %s", err)
	//	return
	//}
	//defer c.Close()
	//
	//client := mtproto.NewRPCAuthKeyClient(c)

	etcdConfg := clientv3.Config{
		Endpoints: discovery.EtcdAddrs,
	}
	r := etcd3.NewResolver("/nebulaim", discovery.ServiceName, etcdConfg)
	var b grpc.Balancer
	switch discovery.Balancer {
	case "random":
		b = load_balancer.NewBalancer(r, load_balancer.NewRandomSelector())
	case "round_robin":
		b = load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	case "consistent_hash":
		b = load_balancer.NewBalancer(r, load_balancer.NewKetamaSelector(load_balancer.DefaultKetamaKey))
	default:
		b = load_balancer.NewBalancer(r, load_balancer.NewRoundRobinSelector())
	}

	c, err = grpc.Dial("", grpc.WithInsecure(),  grpc.WithBalancer(b), grpc.WithTimeout(time.Second*5))
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	return
}

type RPCClient struct {
	conn *grpc.ClientConn
}

func NewRPCClient(target string) (c *RPCClient, err error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	c = &RPCClient{
		conn: conn,
	}
	return
}

// 通用grpc转发器
func (c* RPCClient) Invoke(rpcMetaData *zproto.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	t := mtproto.FindRPCContextTuple(object)
	if t == nil {
		err := fmt.Errorf("Invoke error: %v not regist!\n", object)
		return nil, err
	}

	glog.Infof("Invoke - method: {%s}, req: {%v}", t.Method, object)
	r := t.NewReplyFunc()
	// glog.Infof("Invoke - NewReplyFunc: {%v}\n", r)

	var header, trailer metadata.MD

	// ctx := context.Background()
	ctx, _ := RpcMetadatToOutgoing(context.Background(), rpcMetaData)
	err := c.conn.Invoke(ctx, t.Method, object, r, grpc.Header(&header), grpc.Trailer(&trailer))
	// TODO(@benqi): process header from server
	// grpc.Header(&header)

	if err != nil {
		// TODO(@benqi): 哪些情况需要断开客户端连接
		if s, ok := status.FromError(err); !ok {
			switch s.Code() {
			// TODO(@benqi): Rpc error, trailer has rpc_error metadata
			case codes.Unknown:
				return RpcErrorFromMD(trailer), nil
			}
		}
		glog.Errorf("RPC method: %s,  >> %v.Invoke(_) = _, %v: \n", t.Method, c.conn, err)
		return mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR"), nil
	} else {
		glog.Infof("Invoke - Invoke reply: {%v}\n", r)
		reply, ok := r.(mtproto.TLObject)

		glog.Infof("Invoke %s time: %d", t.Method, (time.Now().Unix() - rpcMetaData.ReceiveTime))

		if !ok {
			err = fmt.Errorf("Invalid reply type, maybe server side bug, %v\n", reply)
			glog.Error(err)
			return mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR"), nil
		}

		return reply, nil
	}
}
