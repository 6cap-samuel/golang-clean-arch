package usecases

import (
	"fmt"
	"golang-clean-arch/entities"
	"golang-clean-arch/usecases/in"
	"golang-clean-arch/usecases/out"
)

type retrievePostInput struct {
	postDataSoruce  out.PostDataSource
	storeDataSource out.StoreDataSource
}

func NewRetrievePostInteractor(
	postDataSoruce *out.PostDataSource,
	storeDataSource *out.StoreDataSource,
) in.RetrievePostsInput {
	return &retrievePostInput{
		*postDataSoruce,
		*storeDataSource,
	}
}

func (r *retrievePostInput) GetAll() (response []entities.Post) {
	var posts = r.postDataSoruce.GetAll()
	var resultList = make([]entities.Post, len(posts)-1)

	for _, item := range posts {
		item.Store = r.storeDataSource.GetStore(item.Store.Id)
		fmt.Println(item)
		resultList = append(resultList, item)
	}

	return resultList
}
