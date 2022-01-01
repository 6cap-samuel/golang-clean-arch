package usecases

import (
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type createPostInput struct {
	postDataSource    out.PostDataSource
	hashtagDataSource out.HashtagDataSource
}

func NewCreatePostInput(
	post *out.PostDataSource,
	hashtag *out.HashtagDataSource,
) in.CreatePostInput {
	return &createPostInput{
		*post,
		*hashtag,
	}
}

func (c *createPostInput) Create(
	post entities.Post,
	hashtags []entities.Hashtag,
) {
	go c.postDataSource.Create(post)
	go c.hashtagDataSource.Create(hashtags)
}
