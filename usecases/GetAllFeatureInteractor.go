package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
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
