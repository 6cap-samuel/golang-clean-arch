package out

import "feature-service/entities"

type FeatureDataSource interface {
	GetAll() (response []entities.Feature)
}
