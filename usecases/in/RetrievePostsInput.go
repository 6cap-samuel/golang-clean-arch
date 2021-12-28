package in

import "golang-clean-arch/entities"

type RetrievePostsInput interface {
	GetAll() (response []entities.Post)
}
