package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/nych/docker-in-a-nutshell/demos/grpc-demo/gen/go/greeter/v1/greeterv1connect"
	"github.com/nych/docker-in-a-nutshell/demos/grpc-demo/internal/greeter"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	serverPortFlag string
)

func init() {
	flag.StringVar(&serverPortFlag, "port", "8081", "grpc server address in format host:port")
	flag.Parse()
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(greeterv1connect.NewGreeterServiceHandler(&greeter.Server{}))
	log.Println("start serving..")
	http.ListenAndServe(fmt.Sprintf(":%s", serverPortFlag), h2c.NewHandler(mux, &http2.Server{}))
}
