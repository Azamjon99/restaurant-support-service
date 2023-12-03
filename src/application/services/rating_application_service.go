package application

import (
	"context"
	"fmt"

	"github.com/Azamjon99/restaurant-support-service/src/application/dtos"
	pb "github.com/Azamjon99/restaurant-support-service/src/application/protos/restaurant_support"
	"github.com/Azamjon99/restaurant-support-service/src/domain/services"
)

type RatingApplicationService interface {
	CreateRating(ctx context.Context, res *pb.Rating) error
	UpdateRating(ctx context.Context, res *pb.Rating) error
	ListRatingByEntity(ctx context.Context, req *pb.ListRatingRequest) (*pb.ListRatingResponse, error)
}

type ratingAppSvcImpl struct {
	ratingSvc services.RatingService
}

func NewRatingApplicationService(ratingSvc services.RatingService) RatingApplicationService {
	return &ratingAppSvcImpl{
		ratingSvc: ratingSvc,
	}
}

func (p *ratingAppSvcImpl) CreateRating(ctx context.Context, req *pb.Rating) error {

	if req.GetEaterId() == "" {
		return fmt.Errorf("Invalid or missing eaterId: %s", req.GetEaterId())
	}

	if req.GetComment() == "" {
		return fmt.Errorf("Invalid or missing Comment: %s", req.GetComment())
	}

	if req.GetRating() == 0 {
		return fmt.Errorf("Invalid or missing Rating: %d", req.GetRating())
	}

	err := p.ratingSvc.CreateRating(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (p *ratingAppSvcImpl) UpdateRating(ctx context.Context, req *pb.Rating) error {

	if req.GetId() == "" {
		return fmt.Errorf("Invalid or missing eaterId: %s", req.GetId())
	}

	if req.GetComment() == "" {
		return fmt.Errorf("Invalid or missing Comment: %s", req.GetComment())
	}

	if req.GetRating() == 0 {
		return fmt.Errorf("Invalid or missing Rating: %d", req.GetRating())
	}

	err := p.ratingSvc.UpdateRating(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (p *ratingAppSvcImpl) ListRatingByEntity(ctx context.Context, req *pb.ListRatingRequest) (*pb.ListRatingResponse, error) {

	if req.GetEntityId() == "" {
		return nil, fmt.Errorf("Invalid or missing GetEntityId: %s", req.GetEntityId())
	}

	if req.GetSort() == "" {
		return nil, fmt.Errorf("Invalid or missing GetSort: %s", req.GetSort())
	}

	if req.GetPage() == 0 {
		return nil, fmt.Errorf("Invalid or missing GetPage: %d", req.GetPage())
	}

	if req.GetPageSize() == 0 {
		return nil, fmt.Errorf("Invalid or missing GetPageSize: %d", req.GetPageSize())
	}

	ratings, err := p.ratingSvc.ListRatingByEntity(ctx, req.GetEntityId(), req.GetSort(), int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, err
	}

	return &pb.ListRatingResponse{
		Ratings: dtos.ToRatingListResponsePB(ratings),
	}, nil
}
