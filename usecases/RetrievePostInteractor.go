package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type retrievePostInput struct {
	postDataSource out.PostDataSource
}

func NewRetrievePostInteractor(
	postDataSoruce *out.PostDataSource,
) in.RetrievePostsInput {
	return &retrievePostInput{
		*postDataSoruce,
	}
}

func (r *retrievePostInput) GetAll() (response []entities.Post) {
	return r.postDataSource.GetAll()
}
