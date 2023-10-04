package main

//go get google.golang.org/grpc  先下载这个东西
//go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  然后下载这个东西
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  继续下载这个东西
//protoc --go_out=. hello.proto
//protoc --go-grpc_out=. hello.proto

import (
	"context"
	"fmt"
	pb "gin_study/gRPCDemo01/hello-server/proto"
	"google.golang.org/grpc"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//在grpc服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}
}
