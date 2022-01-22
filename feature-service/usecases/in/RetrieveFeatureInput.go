package in

import "feature-service/entities"

type RetrieveFeatureInput interface {
	GetAll() (response []entities.Feature)
}
