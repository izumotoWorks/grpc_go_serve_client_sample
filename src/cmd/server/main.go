package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	hellopb "hello/src/grpc/gen"
)

type myServer struct {
	hellopb.UnimplementedGreeterServer
}

func (s *myServer) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	// リクエストからnameフィールドを取り出して
	// "Hello, [名前]!"という文字列を返す
	return &hellopb.HelloReply {
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}


func main() {
	// 1. 8080番ポートのListenerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバを作成
	s := grpc.NewServer()

	// 2.1 gRPCサーバにGreeterサービスを登録
	hellopb.RegisterGreeterServer(s, NewMyServer())

	// 2.2 gRPCサーバにリフレクションサービスを登録
	reflection.Register(s)

	// 3. 作成したgRPCサーバを8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC servr port: %v", port)
		s.Serve(listener)
	}()

	// 4. サーバの終了を待機
	// このサンプルではCtrl+Cで終了するまで待機
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}