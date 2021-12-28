package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

// init concrete class
type createPostInput struct {
	source out.PostDataSource
}

func NewCreatePostInput(
	source *out.PostDataSource,
) in.CreatePostInput {
	return &createPostInput{
		*source,
	}
}

func (c *createPostInput) Create(post entities.Post) {
	c.source.Create(post)
}
