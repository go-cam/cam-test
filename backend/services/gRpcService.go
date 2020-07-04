package services

import (
	"context"
	pb "test/backend-grpc"
)

type gRpcService struct {
	pb.UnimplementedGrpcServer
}

func NewGRpcService() *gRpcService {
	return &gRpcService{}
}

func (srv *gRpcService) Hello(ctx context.Context, recv *pb.Grpc_Hello_Recv) (*pb.Grpc_Hello_Send, error) {
	return &pb.Grpc_Hello_Send{Name: recv.Name}, nil
}
