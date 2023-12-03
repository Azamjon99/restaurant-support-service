package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-support-service/src/domain/models"
	"github.com/Azamjon99/restaurant-support-service/src/domain/repositories"
	"gorm.io/gorm"
)

const (
	tableRating = "rating.ratings"
)

type RatingRespository interface {
	SaveRating(ctx context.Context, rating *models.Rating) error
	UpdateRating(ctx context.Context, rating *models.Rating) error
	ListRatingsByEntity(ctx context.Context, entityID, sort string, page, pageSize int) ([]*models.Rating, error)
}

type ratingRepoImpl struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) repositories.RatingRepository {
	return &ratingRepoImpl{
		db: db,
	}
}

func (a *ratingRepoImpl) SaveRating(ctx context.Context, rating *models.Rating) error {
	res := a.db.WithContext(ctx).Table(tableRating).Create(rating)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *ratingRepoImpl) UpdateRating(ctx context.Context, rating *models.Rating) error {

	res := r.db.WithContext(ctx).Table(tableRating).Save(rating)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (a *ratingRepoImpl) ListRatingsByEntity(ctx context.Context, entityID, sort string, page, pageSize int) ([]*models.Rating, error) {

	var ratings []*models.Rating

	res := a.db.WithContext(ctx).Table(tableRating).Where("eater_id = ?", entityID).Find(ratings)

	if res.Error != nil {
		return nil, res.Error
	}

	return ratings, nil
}
