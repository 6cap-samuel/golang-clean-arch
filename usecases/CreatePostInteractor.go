package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type createPostInput struct {
	postDataSource out.PostDataSource
}

func NewCreatePostInput(
	post *out.PostDataSource,
) in.CreatePostInput {
	return &createPostInput{
		*post,
	}
}

func (c *createPostInput) Create(post entities.Post) {
	c.postDataSource.Create(post)
}
