package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type updatePostInput struct {
	postDataSource out.PostDataSource
}

func NewUpdatePostIntractor(
	postDataSoruce *out.PostDataSource,
) in.UpdatePostInput {
	return &updatePostInput{
		*postDataSoruce,
	}
}

func (u updatePostInput) UpdateFood(postId string, food []entities.Food) {
	u.postDataSource.UpdateFood(postId, food)
}
