// Copyright (c) 2022 zhaochun
// gmgo is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

/*
grpc_test 是对`gitee.com/zhaochuninhefei/gmgo/grpc`的测试包
*/

package grpc_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"testing"
	"time"

	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"github.com/hxx258456/ccgo/gmtls"
	"github.com/hxx258456/ccgo/grpc"
	"github.com/hxx258456/ccgo/grpc/credentials"
	"github.com/hxx258456/ccgo/grpc/grpc_test/echo"
	"github.com/hxx258456/ccgo/net/context"
	"github.com/hxx258456/ccgo/x509"
)

const (
	port     = ":50051"
	address  = "localhost:50051"
	ca       = "testdata/ca.cert"
	signCert = "testdata/sign.cert"
	signKey  = "testdata/sign.key"
	userCert = "testdata/user.cert"
	userKey  = "testdata/user.key"
)

var end chan bool

func Test_credentials(t *testing.T) {
	zcgologConfig := &zclog.Config{
		LogLevelGlobal: zclog.LOG_LEVEL_DEBUG,
	}
	zclog.InitLogger(zcgologConfig)
	end = make(chan bool, 64)
	go serverRun()
	time.Sleep(1000000)
	go clientRun()
	<-end
}

func serverRun() {
	signCert, err := gmtls.LoadX509KeyPair(signCert, signKey)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	cacert, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(cacert)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	creds := credentials.NewTLS(&gmtls.Config{
		ClientAuth:   gmtls.RequireAndVerifyClientCert,
		Certificates: []gmtls.Certificate{signCert},
		ClientCAs:    certPool,
	})
	s := grpc.NewServer(grpc.Creds(creds))
	echo.RegisterEchoServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Serve: %v", err)
	}
}

func clientRun() {
	cert, err := gmtls.LoadX509KeyPair(userCert, userKey)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	cacert, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(cacert)
	creds := credentials.NewTLS(&gmtls.Config{
		ServerName:   "server.test.com",
		Certificates: []gmtls.Certificate{cert},
		RootCAs:      certPool,
		ClientAuth:   gmtls.RequireAndVerifyClientCert,
	})
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("cannot to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	c := echo.NewEchoClient(conn)
	echoInClient(c)
	end <- true
}

// 客户端echo处理
func echoInClient(c echo.EchoClient) {
	msgClient := "hello, this is client."
	fmt.Printf("客户端发出消息: %s\n", msgClient)
	r, err := c.Echo(context.Background(), &echo.EchoRequest{Req: msgClient})
	if err != nil {
		log.Fatalf("failed to echo: %v", err)
	}
	msgServer := r.Result
	fmt.Printf("客户端收到消息: %s\n", msgServer)
}

type server struct{}

// Echo 服务端echo处理
//
//goland:noinspection GoUnusedParameter
func (s *server) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	msgClient := req.Req
	fmt.Printf("服务端接收到消息: %s\n", msgClient)
	msgServer := "hello,this is server."
	fmt.Printf("服务端返回消息: %s\n", msgServer)
	return &echo.EchoResponse{Result: msgServer}, nil
}
