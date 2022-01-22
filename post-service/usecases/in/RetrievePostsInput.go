package in

import "golang-clean-arch/entities"

type RetrievePostsInput interface {
	GetAll(filters []string) (response []entities.Post)
}
