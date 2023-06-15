package greeter

import (
	"context"
	"fmt"
	"log"

	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/nych/docker-in-a-nutshell/demos/grpc-demo/gen/go/greeter/v1"
)

type Server struct{}

func (s *Server) GetGreetings(ctx context.Context, request *connect_go.Request[v1.GetGreetingsRequest]) (*connect_go.Response[v1.GetGreetingsResponse], error) {
	log.Println("processing GetGreetings request..")
	return connect_go.NewResponse(&v1.GetGreetingsResponse{
		Greetings: fmt.Sprintf("Hello dear %s. I wish you a wonderful day!", request.Msg.GetTo()),
	}), nil
}
