package in

import "golang-clean-arch/entities"

type UpdatePostInput interface {
	UpdateFood(postId string, food []entities.Food)
}
