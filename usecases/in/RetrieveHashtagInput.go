package in

import "golang-clean-arch/entities"

type RetrieveHashtagInput interface {
	GetAll() (response []entities.Hashtag)
}
