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

func (r *retrievePostInput) GetAll(filters []string) (response []entities.Post) {
	if len(filters) == 0 {
		return r.postDataSource.GetAll()
	} else {
		return r.postDataSource.GetAllWith(filters)
	}
}
