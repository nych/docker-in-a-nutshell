package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	greeterv1 "github.com/nych/docker-in-a-nutshell/demos/grpc-demo/gen/go/greeter/v1"
	"github.com/nych/docker-in-a-nutshell/demos/grpc-demo/internal/greeter"
	"google.golang.org/grpc"
)

var (
	serverPortFlag string
)

func init() {
	flag.StringVar(&serverPortFlag, "port", "8081", "grpc server address in format host:port")
	flag.Parse()
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPortFlag))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	greeterv1.RegisterGreeterServiceServer(srv, &greeter.Server{})

	log.Println("start serving..")
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
