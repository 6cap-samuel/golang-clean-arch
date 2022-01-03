package out

import "golang-clean-arch/entities"

type PostDataSource interface {
	Get(postId string) (response entities.Post)
	GetAll() (response []entities.Post)
	GetAllWith(filters []string) (response []entities.Post)
	Create(post entities.Post)
	UpdateFood(postId string, food []entities.Food)
}
