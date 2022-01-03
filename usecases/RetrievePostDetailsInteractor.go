package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type retrievePostDetailsInteractor struct {
	postDataSource out.PostDataSource
}

func NewRetrievePostDetailsInteractor(
	postDataSource *out.PostDataSource,
) in.RetrievePostDetailsInput {
	return &retrievePostDetailsInteractor{
		*postDataSource,
	}
}

func (r retrievePostDetailsInteractor) Get(
	postId string,
) (response entities.Post) {
	return r.postDataSource.Get(postId)
}
