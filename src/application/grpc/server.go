package grpc

import (
	"context"

	pb "github.com/Azamjon99/restaurant-support-service/src/application/protos/restaurant_support"

	application "github.com/Azamjon99/restaurant-support-service/src/application/services"
)

type Server struct {
	pb.SupportServiceServer
	ratingApp application.RatingApplicationService
}

func NewServer(
	ratingApp application.RatingApplicationService,
) *Server {
	return &Server{
		ratingApp: ratingApp,
	}
}

func (s *Server) CreateRating(ctx context.Context, res *pb.Rating) error {
	return s.CreateRating(ctx, res)
}
