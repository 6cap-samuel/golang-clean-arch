package in

import "golang-clean-arch/entities"

type RetrievePostDetailsInput interface {
	Get(postId string) (response entities.Post)
}
