package main

import (
	"log"
	"net"

	"github.com/koddr/example-go-grpc-server/pkg/adder"
	"github.com/koddr/example-go-grpc-server/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	// Register gRPC server
	api.RegisterAdderServer(s, srv)

	// Listen on port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
