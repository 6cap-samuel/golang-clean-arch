package mappers

import (
	"golang-clean-arch/configurations"
	"golang-clean-arch/controllers/requests"
	"golang-clean-arch/entities"
)

func UpdateFoodRequestToFoodMapper(requests requests.UpdateFoodRequest) []entities.Food {
	var response []entities.Food

	for _, foodRequest := range requests.Foods {
		response = append(response, entities.Food{
			Id:     configurations.NewIdentity(),
			Name:   foodRequest.Name,
			Image:  foodRequest.Image,
			Cost:   foodRequest.Cost,
			Rating: foodRequest.Rating,
		})
	}

	return response
}
