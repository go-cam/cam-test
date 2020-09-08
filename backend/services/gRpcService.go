package services

import (
	"context"
	pb "test/backend-grpc"
)

type HelloWorldService struct {
	pb.UnimplementedHelloWorldServer
}

func (srv *HelloWorldService) SayHello(ctx context.Context, recv *pb.HelloWorld_SayHello_Recv) (*pb.HelloWorld_SayHello_Send, error) {
	return &pb.HelloWorld_SayHello_Send{Name: recv.Name}, nil
}
