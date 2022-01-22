package in

import "golang-clean-arch/entities"

type RetrieveFeatureInput interface {
	GetAll() (response []entities.Feature)
}
