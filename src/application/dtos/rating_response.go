package dtos

import (
	pb "github.com/Azamjon99/restaurant-support-service/src/application/protos/restaurant_support"
	"github.com/Azamjon99/restaurant-support-service/src/domain/models"
)

func ToRatingPB(rating *models.Rating) *pb.Rating {
	return &pb.Rating{
		Id:       rating.ID,
		EntityId: rating.EntityID,
		EaterId:  rating.EaterID,
	}
}

func ToRatingListResponsePB(rating []*models.Rating) []*pb.Rating {
	ratingArr := make([]*pb.Rating, len(rating))
	for i, r := range rating {
		ratingArr[i] = ToRatingPB(r)
	}
	return ratingArr
}
