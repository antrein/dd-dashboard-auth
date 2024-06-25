package auth

import (
	"context"

	pb "github.com/antrein/proto-repository/pb/dd"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return nil, nil
}
