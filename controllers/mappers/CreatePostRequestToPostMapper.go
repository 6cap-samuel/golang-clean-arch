package mappers

import (
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers/requests"
	"golang-clean-arch/entities"
)

func CreatePostRequestToPostMapper(request requests.CreatePostRequest) *entities.Post {

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
		Rating: request.Rating,
	}
}
