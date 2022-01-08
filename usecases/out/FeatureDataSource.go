package out

import "golang-clean-arch/entities"

type FeatureDataSource interface {
	GetAll() (response []entities.Feature)
}
