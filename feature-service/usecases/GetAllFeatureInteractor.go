package usecases

import (
	"feature-service/entities"
	"feature-service/usecases/in"
	"feature-service/usecases/out"
)

type getAllFeatureInteractor struct {
	feature out.FeatureDataSource
}

func NewGetAllFeatureDataInteractor(
	feature *out.FeatureDataSource,
) in.RetrieveFeatureInput {
	return &getAllFeatureInteractor{
		*feature,
	}
}

func (g getAllFeatureInteractor) GetAll() (
	response []entities.Feature,
) {
	return g.feature.GetAll()
}
