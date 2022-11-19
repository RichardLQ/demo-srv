package main

import (
	"github.com/RichardLQ/user-srv/proto/stub"
	"google.golang.org/grpc"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	stub.RegisterUserServiceServer(rpcServer, stub.UserService)
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}
