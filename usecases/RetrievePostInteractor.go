package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

// init concrete class
type retrievePostInput struct {
	source out.PostDataSource
}

func NewRetrievePostInteractor(
	source *out.PostDataSource,
) in.RetrievePostsInput {
	return &retrievePostInput{
		*source,
	}
}

func (r *retrievePostInput) GetAll() (response []entities.Post) {
	return r.source.GetAll()
}
