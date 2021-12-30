package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-clean-arch/configurations"
	"golang-clean-arch/entities"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/usecases/out"
)

type postRepository struct {
	Collection *mongo.Collection
}

func NewPostRepository(database *mongo.Database) out.PostDataSource {
	return &postRepository{
		Collection: database.Collection("posts"),
	}
}

func (p postRepository) GetAll() (response []entities.Post) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	cursor, err := p.Collection.Find(ctx, bson.M{})
	exceptions.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exceptions.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, entities.Post{
			Id:          document["_id"].(string),
			Description: document["description"].(string),
			Store: entities.Store{
				Id: document["store"].(string),
			},
		})
	}
	return response
}

func (p postRepository) Create(post entities.Post) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	_, err := p.Collection.InsertOne(ctx, bson.M{
		"_id":         post.Id,
		"description": post.Description,
		"store":       post.Store.Id,
	})

	exceptions.PanicIfNeeded(err)
}
