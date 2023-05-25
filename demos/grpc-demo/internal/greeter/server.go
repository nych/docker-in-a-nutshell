package greeter

import (
	"context"
	"fmt"
	"log"

	greeterv1 "github.com/nych/docker-in-a-nutshell/demos/grpc-demo/gen/go/greeter/v1"
)

type Server struct {
	greeterv1.UnimplementedGreeterServiceServer
}

func (s *Server) GetGreetings(ctx context.Context, request *greeterv1.GetGreetingsRequest) (*greeterv1.GetGreetingsResponse, error) {
	log.Println("processing GetGreetings request..")
	return &greeterv1.GetGreetingsResponse{
		Greetings: fmt.Sprintf("Hello dear %s. I wish you a wonderful day!", request.GetTo()),
	}, nil
}
