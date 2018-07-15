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
	"sync"
	"encoding/binary"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/proto/zproto"
	"github.com/nebulaim/telegramd/baselib/bytes2"
)

type sessionManager struct {
	sessions sync.Map // map[int64]*sessionClientList
}

func newSessionManager() *sessionManager {
	return &sessionManager{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *sessionManager) onSessionData(clientConnID uint64, md *zproto.ZProtoMetadata, sessData *zproto.ZProtoSessionData) error {
	glog.Infof("onSessionData: data: {client_conn_id: %d, frontendConnID: %d, connType: %d, md: %v, buf_len: %d, buf: %s}",
		// conn.RemoteAddr(),
		clientConnID,
		sessData.SessionId,
		md,
		len(sessData.MtpRawData),
		bytes2.Dump(sessData.MtpRawData))

	////
	authKeyId := int64(binary.LittleEndian.Uint64(sessData.MtpRawData))

	// TODO(@benqi): sync s.sessions
	var sessList *clientSessionManager
	if vv, ok := s.sessions.Load(authKeyId); !ok {
		authKey := getCacheAuthKey(authKeyId)
		if authKey == nil {
			err := fmt.Errorf("onSessionData - not found authKeyId: {%d}", authKeyId)
			glog.Error(err)
			return err
		}

		sessList = newClientSessionManager(authKeyId, authKey, 0)
		s.sessions.Store(authKeyId, sessList)
		s.onNewSessionClientManager(sessList)
	} else {
		sessList, _ = vv.(*clientSessionManager)
	}

	return sessList.OnSessionDataArrived(makeClientConnID(sessData.ConnType, clientConnID, sessData.SessionId), md, sessData.MtpRawData)
}

func (s *sessionManager) onSyncData(authKeyId, sessionId int64, md *zproto.ZProtoMetadata, data *messageData) error {
	var sessList *clientSessionManager

	if vv, ok := s.sessions.Load(authKeyId); !ok {
		err := fmt.Errorf("pushToSessionData - not find sessionList by authKeyId: {%d}", authKeyId)
		glog.Warning(err)
		return err
	} else {
		sessList, _ = vv.(*clientSessionManager)
	}

	return sessList.OnSyncDataArrived(sessionId, md, data)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
// session event
func (s *sessionManager) onNewSessionClientManager(sess *clientSessionManager) {
	sess.Start()
}

func (s *sessionManager) onCloseSessionClientManager(authKeyId int64) {
	if vv, ok := s.sessions.Load(authKeyId); ok {
		vv.(*clientSessionManager).Stop()
		s.sessions.Delete(authKeyId)
	}
}
