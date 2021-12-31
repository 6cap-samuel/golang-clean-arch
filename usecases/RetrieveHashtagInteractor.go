package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type retrieveHashtagInteractor struct {
	hashtagDataSource out.HashtagDataSource
}

func NewRetrieveHashtagInteractor(
	hashtagDataSource *out.HashtagDataSource,
) in.RetrieveHashtagInput {
	return &retrieveHashtagInteractor{
		*hashtagDataSource,
	}
}

func (r retrieveHashtagInteractor) GetAll() (response []entities.Hashtag) {
	return r.hashtagDataSource.GetAll()
}
