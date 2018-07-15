/*
 *  Copyright (c) 2018, https://github.com/nebulaim
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

package server

import (
	"fmt"
	"time"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/grpc_util"
	"github.com/nebulaim/telegramd/baselib/net2"
	"github.com/nebulaim/telegramd/baselib/redis_client"
	"github.com/nebulaim/telegramd/biz/dal/dao"
	"github.com/nebulaim/telegramd/service/idgen/client"
	"github.com/nebulaim/telegramd/baselib/base"
	"github.com/nebulaim/telegramd/proto/zproto"
	"github.com/nebulaim/telegramd/proto/mtproto"
)

type SessionServer struct {
	idgen      	   idgen.UUIDGen
	server         *zproto.ZProtoServer
	bizRpcClient   *grpc_util.RPCClient
	nbfsRpcClient  *grpc_util.RPCClient
	syncRpcClient  mtproto.RPCSyncClient
	sessionManager *sessionManager
	syncHandler    *syncHandler
}

func NewSessionServer() *SessionServer {
	return &SessionServer{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// AppInstance interface
func (s *SessionServer) Initialize() error {
	err := InitializeConfig()
	if err != nil {
		glog.Fatal(err)
		return err
	}
	glog.Info("load conf: ", Conf)

	// idgen
	s.idgen, _ = idgen.NewUUIDGen("snowflake", base.Int32ToString(Conf.ServerId))

	// 初始化mysql_client、redis_client
	redis_client.InstallRedisClientManager(Conf.Redis)

	// 初始化redis_dao、mysql_dao
	dao.InstallRedisDAOManager(redis_client.GetRedisClientManager())
	// TODO(@benqi): config cap
	InitCacheAuthManager(1024*1024, &Conf.AuthKeyRpcClient)

	s.sessionManager = newSessionManager()
	s.syncHandler = newSyncHandler(s.sessionManager)
	s.server = zproto.NewZProtoServer(Conf.Server, s)

	return nil
}

func (s *SessionServer) RunLoop() {
	// TODO(@benqi): check error
	// timingWheel.Start()

	s.bizRpcClient, _ = grpc_util.NewRPCClient(&Conf.BizRpcClient)
	s.nbfsRpcClient, _ = grpc_util.NewRPCClient(&Conf.NbfsRpcClient)
	c, _ := grpc_util.NewRPCClient(&Conf.SyncRpcClient)
	s.syncRpcClient = mtproto.NewRPCSyncClient(c.GetClientConn())

	s.server.Serve()
}

func (s *SessionServer) Destroy() {
	glog.Infof("sessionServer - destroy...")
	s.server.Stop()
	time.Sleep(1 * time.Second)
	// s.client.Stop()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// TcpConnectionCallback
func (s *SessionServer) OnServerNewConnection(conn *net2.TcpConnection) {
	glog.Infof("OnNewConnection %v", conn.RemoteAddr())
}

func (s *SessionServer) OnServerMessageDataArrived(conn *net2.TcpConnection, md *zproto.ZProtoMetadata, sessionId, messageId uint64, seqNo uint32, msg zproto.MessageBase) error {
	switch msg.(type) {
	case *zproto.ZProtoSessionClientNew:
		glog.Info("onServerMessageDataArrived - sessionClientNew: ", conn)
		return nil
	case *zproto.ZProtoSessionData:
		return s.sessionManager.onSessionData(conn.GetConnID(), md, msg.(*zproto.ZProtoSessionData))
	case *zproto.ZProtoSyncData:
		sres, err := s.syncHandler.onSyncData(conn, msg.(*zproto.ZProtoSyncData))
		if err != nil {
			glog.Error(err)
			return nil
		}
		return zproto.SendMessageByConn(conn, md, sres)
	case *zproto.ZProtoSessionClientClosed:
		glog.Info("onServerMessageDataArrived - sessionClientClosed: ", conn)
		return nil
	default:
		err := fmt.Errorf("invalid payload type: %v", msg)
		glog.Error(err)
		return err
	}
}

func (s *SessionServer) OnServerConnectionClosed(conn *net2.TcpConnection) {
	glog.Infof("OnConnectionClosed - %v", conn.RemoteAddr())
}

//func (s *SessionServer) SendToClientData(connID, sessionID uint64, md *zproto.ZProtoMetadata, buf []byte) error {
//	glog.Infof("sendToClientData - {%d, %d}", connID, sessionID)
//	//conn := s.server.GetConnection(connID)
//	//if conn != nil {
//	//	return sendDataByConnection(conn, sessionID, md, buf)
//	//} else {
//	//	err := fmt.Errorf("send data error, conn offline, connID: %d", connID)
//	//	glog.Error(err)
//	//	return err
//	//}
//	return nil
//}
