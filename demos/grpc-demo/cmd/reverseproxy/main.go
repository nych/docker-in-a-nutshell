package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/nych/docker-in-a-nutshell/demos/grpc-demo/gen/go/greeter/v1"
)

var (
	grpcServerAddressFlag string
	gatewayPortFlag       string
)

func init() {
	flag.StringVar(&grpcServerAddressFlag, "grpcServer", "localhost:8081", "grpc server address in format host:port")
	flag.StringVar(&gatewayPortFlag, "port", "8080", "http gateway port")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithNoProxy(),
	}
	if err := gw.RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, grpcServerAddressFlag, opts); err != nil {
		panic(err)
	}
	log.Println("registered handler..")

	http.ListenAndServe(fmt.Sprintf(":%s", gatewayPortFlag), mux)
}
