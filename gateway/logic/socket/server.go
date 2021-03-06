// Copyright 2018 github.com/xiaoenai. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package socket

import (
	"net"

	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/plugin/auth"
	"github.com/henrylee2cn/erpc/v6/plugin/proxy"
	micro "github.com/xiaoenai/tp-micro/v6"
	"github.com/xiaoenai/tp-micro/v6/clientele"
	"github.com/xiaoenai/tp-micro/v6/discovery"
	"github.com/xiaoenai/tp-micro/v6/gateway/logic"
	websocket "github.com/xiaoenai/tp-micro/v6/gateway/logic/websocket"
)

var (
	outerPeer      erpc.Peer
	outerServer    *micro.Server
	clientAuthInfo string
)

// OuterServeConn serves connetion.
func OuterServeConn(conn net.Conn) {
	sess, err := outerServer.ServeConn(conn)
	if err != nil {
		erpc.Errorf("Serve net.Conn error: %v", err)
	}
	<-sess.CloseNotify()
}

// Serve starts TCP gateway service.
func Serve(outerSrvCfg, innerSrvCfg micro.SrvConfig, protoFunc erpc.ProtoFunc) {
	outerServer = micro.NewServer(
		outerSrvCfg,
		authChecker,
		socketConnTabPlugin,
		proxy.NewPlugin(logic.ProxySelector),
		preWritePushPlugin(),
	)

	outerPeer = outerServer.Peer()

	innerPlugins := logic.InnerServerPlugins()
	discoveryService := discovery.ServicePluginFromEtcd(
		innerSrvCfg.InnerIpPort(),
		clientele.GetEtcdClient(),
	)
	innerPlugins = append(innerPlugins, discoveryService)
	innerServer := micro.NewServer(
		innerSrvCfg,
		innerPlugins...,
	)

	gwGroup := innerServer.SubRoute("/gw")
	{
		verGroup := gwGroup.SubRoute(logic.ApiVersion())
		{
			verGroup.RouteCallFunc((*gw).Hosts)
			// socket api
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*gw).SocketTotal))
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*gw).SocketPush))
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*gw).SocketMpush))
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*gw).SocketKick))
			// ws api
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*websocket.Gw).WsTotal))
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*websocket.Gw).WsPush))
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*websocket.Gw).WsMpush))
			discoveryService.ExcludeApi(verGroup.RouteCallFunc((*websocket.Gw).WsKick))
		}
	}

	go outerServer.ListenAndServe(protoFunc)
	go innerServer.ListenAndServe(protoFunc)

	select {}
}

// auth plugin
var authChecker = auth.NewCheckerPlugin(
	func(sess auth.Session, fn auth.RecvOnce) (ret interface{}, stat *erpc.Status) {
		var authInfo string
		stat = fn(&authInfo)
		if !stat.OK() {
			return
		}
		erpc.Tracef("auth info: %v", authInfo)
		stat = socketConnTabPlugin.authAndLogon(authInfo, sess)
		if !stat.OK() {
			return
		}
		return "", nil
	},
	erpc.WithBodyCodec('s'),
)

// preWritePushPlugin returns PreWritePushPlugin.
func preWritePushPlugin() erpc.Plugin {
	return &perPusher{fn: logic.SocketHooks().PreWritePush}
}

type perPusher struct {
	fn func(erpc.WriteCtx) *erpc.Status
}

func (p *perPusher) Name() string {
	return "PUSH-LOGIC"
}

var (
	_ erpc.PreWritePushPlugin = (*perPusher)(nil)
)

func (p *perPusher) PreWritePush(ctx erpc.WriteCtx) *erpc.Status {
	return p.fn(ctx)
}
