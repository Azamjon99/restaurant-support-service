package services

import (
	"context"
	"time"

	pb "github.com/Azamjon99/restaurant-support-service/src/application/protos/restaurant_support"
	"github.com/Azamjon99/restaurant-support-service/src/domain/models"
	"github.com/Azamjon99/restaurant-support-service/src/infrastructure/repositories"
)

type RatingService interface {
	CreateRating(ctx context.Context, req *pb.Rating) error
	UpdateRating(ctx context.Context, req *pb.Rating) error
	ListRatingByEntity(ctx context.Context, entityID, sort string, page, pageSize int) ([]*models.Rating, error)
}

type ratingServiceImpl struct {
	ratingRespo repositories.RatingRespository
}

func NewRestaturantRepoService(ratingRespo repositories.RatingRespository) RatingService {

	return &ratingServiceImpl{
		ratingRespo: ratingRespo,
	}

}

func (r *ratingServiceImpl) CreateRating(ctx context.Context, req *pb.Rating) error {
	eater := &models.EaterInfo{
		ID:   req.GetEaterId(),
		Name: req.GetEater().Name,
	}

	restaturant := &models.RestaurantInfo{
		Name: req.Restaurant.GetName(),
	}

	rating_model := models.Rating{
		ID:         req.GetId(),
		EntityID:   req.GetEntityId(),
		EaterID:    req.GetEaterId(),
		Eater:      eater,
		Restaurant: restaturant,
		Comment:    req.GetComment(),
		Rating:     int(req.GetRating()),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := r.ratingRespo.SaveRating(ctx, &rating_model)

	if err != nil {
		return err
	}

	return nil
}

func (r *ratingServiceImpl) UpdateRating(ctx context.Context, req *pb.Rating) error {
	rating_model := models.Rating{
		ID:        req.GetId(),
		Comment:   req.GetComment(),
		UpdatedAt: time.Now(),
	}

	err := r.ratingRespo.UpdateRating(ctx, &rating_model)

	if err != nil {
		return err
	}

	return nil
}

func (r *ratingServiceImpl) ListRatingByEntity(ctx context.Context, entityID, sort string, page, pageSize int) ([]*models.Rating, error) {

	ratings, err := r.ratingRespo.ListRatingsByEntity(ctx, entityID, sort, page, pageSize)

	if err != nil {
		return nil, err
	}

	return ratings, nil
}
