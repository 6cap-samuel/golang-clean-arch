package out

import "golang-clean-arch/entities"

type HashtagDataSource interface {
	GetAll() (response []entities.Hashtag)
	Create(post []entities.Hashtag)
}
