package main

import (
	"context"
	"fmt"
	pb "gin_study/gRPCDemo01/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// 因为是客户端，所以直接写代码进行连接就可以了
func main() {

	//连接到server端，此处禁用安全传输，没有进行加密和认证
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials())) //前面传入链接地址，就是之前服务端的端口  后面传入加密方式   现在先不使用加密方式
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//标准的两件套

	//建立连接
	client := pb.NewSayHelloClient(conn)

	//执行rpc调用
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "woc youdiankanbumingbai a "})
	fmt.Println(resp.GetResponseMsg())
}
