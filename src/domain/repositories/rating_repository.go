package repositories

import (
	"context"

	"github.com/Azamjon99/restaurant-support-service/src/domain/models"
)

type RatingRepository interface {
	SaveRating(ctx context.Context, rating *models.Rating) error
	UpdateRating(ctx context.Context, rating *models.Rating) error
	ListRatingsByEntity(ctx context.Context, entityID, sort string, page, pageSize int) ([]*models.Rating, error)
}
