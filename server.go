package main

import (
	"gRPC_test/uniqId"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen server on 9000 :%v", err)
	}

	s := uniqId.Server{}

	grpcServer := grpc.NewServer()

	uniqId.RegisterUniqIdServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("Failed to serve gRPC on 9000 :%v", err)
	}
}
