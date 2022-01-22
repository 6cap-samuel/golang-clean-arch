package in

import "golang-clean-arch/entities"

type CreatePostInput interface {
	Create(post entities.Post, hashtags []entities.Hashtag)
}
