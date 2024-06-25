package auth

import (
	"antrein/dd-dashboard-auth/model/config"
	"context"
	"errors"
	"strings"

	pb "github.com/antrein/proto-repository/pb/dd"
	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	resp := &pb.ValidateTokenResponse{
		IsValid: false,
		UserId:  "",
	}
	tokenString := strings.TrimSpace(strings.TrimPrefix(in.Token, "Bearer"))
	if tokenString == "" {
		return resp, errors.New("Invalid token format")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token format")
		}
		return []byte(s.cfg.Secrets.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return resp, errors.New("Token invalid")

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return resp, errors.New("Token invalid")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return resp, errors.New("Token invalid")
	}

	resp.IsValid = true
	resp.UserId = userID

	return resp, nil
}
