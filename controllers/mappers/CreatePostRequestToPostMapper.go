package mappers

import (
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers/requests"
	"golang-clean-arch/entities"
)

func CreatePostRequestToPostMapper(request requests.CreatePostRequest) *entities.Post {
	var ratings []entities.Rating

	for _, item := range request.Ratings {
		ratings = append(ratings, entities.Rating{
			Id:    configurations.NewIdentity(),
			Name:  item.Name,
			Value: item.Value,
		})
	}

	return &entities.Post{
		Id:          configurations.NewIdentity(),
		Description: request.Description,
		Store: entities.Store{
			Id:       configurations.NewIdentity(),
			Name:     request.StoreName,
			Image:    request.StoreImage,
			Location: request.StoreLocation,
			Lat:      request.StoreLat,
			Long:     request.StoreLong,
		},
		Ratings: ratings,
	}
}
