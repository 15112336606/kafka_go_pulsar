// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package network

import (
	"github.com/paashzj/kafka_go_pulsar/pkg/network/ctx"
	"github.com/panjf2000/gnet"
	"github.com/protocol-laboratory/kafka-codec-go/codec"
	"github.com/sirupsen/logrus"
)

func (s *Server) ReactJoinGroup(ctx *ctx.NetworkContext, req *codec.JoinGroupReq) (*codec.JoinGroupResp, gnet.Action) {
	if !s.checkSaslGroup(ctx, req.GroupId) {
		return nil, gnet.Close
	}
	logrus.Debug("join group req", req)
	lowResp, err := s.kafsarImpl.GroupJoin(ctx.Addr, req)
	if err != nil {
		return nil, gnet.Close
	}
	resp := &codec.JoinGroupResp{
		BaseResp: codec.BaseResp{
			CorrelationId: req.CorrelationId,
		},
		ErrorCode:    lowResp.ErrorCode,
		GenerationId: lowResp.GenerationId,
		ProtocolType: lowResp.ProtocolType,
		ProtocolName: lowResp.ProtocolName,
		LeaderId:     lowResp.LeaderId,
		MemberId:     lowResp.MemberId,
		Members:      lowResp.Members,
	}

	logrus.Debug("resp ", resp)
	return resp, gnet.None
}
