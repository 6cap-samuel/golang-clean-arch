package mappers

import (
	"meeting-service/configurations"
	"meeting-service/controllers/requests"
	"meeting-service/entities"
)

func CreatePostRequestToPostMapper(request requests.CreateMeetingRequest) *entities.Meeting {
	return &entities.Meeting{
		Id: configurations.NewIdentity(),
		Store: entities.Store{
			Id:       configurations.NewIdentity(),
			Name:     request.StoreName,
			Image:    request.StoreImage,
			Location: request.StoreLocation,
			Lat:      request.StoreLat,
			Long:     request.StoreLong,
		},
		MenusImageUrl:   request.StoreMenuImages,
		MeetingDateTime: request.MeetingDateTime,
	}
}
