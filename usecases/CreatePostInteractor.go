package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type createPostInput struct {
	postDataSource  out.PostDataSource
	storeDataSource out.StoreDataSource
}

func NewCreatePostInput(
	post *out.PostDataSource,
	store *out.StoreDataSource,
) in.CreatePostInput {
	return &createPostInput{
		*post,
		*store,
	}
}

func (c *createPostInput) Create(post entities.Post) {
	c.storeDataSource.Create(post.Store)
	c.postDataSource.Create(post)
}
