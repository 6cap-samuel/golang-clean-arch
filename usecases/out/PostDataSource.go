package out

import "golang-clean-arch/entities"

type PostDataSource interface {
	GetAll() (response []entities.Post)
	Create(post entities.Post)
	UpdateFood(postId string, food []entities.Food)
}
