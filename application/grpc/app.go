package grpc

import (
	"antrein/dd-dashboard-auth/application/common/resource"
	"antrein/dd-dashboard-auth/application/common/usecase"
	"antrein/dd-dashboard-auth/model/config"
	"context"

	pb "github.com/antrein/proto-repository/pb/dd"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func ApplicationDelegate(cfg *config.Config, uc *usecase.CommonUsecase, rsc *resource.CommonResource) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()

	// Hello service
	helloServer := &helloServer{}
	pb.RegisterGreeterServer(grpcServer, helloServer)

	return grpcServer, nil
}
