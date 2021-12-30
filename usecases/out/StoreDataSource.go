package out

import "golang-clean-arch/entities"

type StoreDataSource interface {
	GetStore(id string) (response entities.Store)
	Create(post entities.Store)
}
