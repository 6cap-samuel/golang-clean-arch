package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-clean-arch/configurations"
	"golang-clean-arch/entities"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/usecases/out"
)

type hashtagRepository struct {
	Collection *mongo.Collection
}

func NewHashtagRepository(database *mongo.Database) out.HashtagDataSource {
	return &hashtagRepository{
		Collection: database.Collection("hashtags"),
	}
}

func (h hashtagRepository) Create(hashtags []entities.Hashtag) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	var insertionInterface []interface{}
	for _, t := range hashtags {
		insertionInterface = append(insertionInterface, t)
	}

	_, err := h.Collection.InsertMany(ctx, insertionInterface)

	exceptions.PanicIfNeeded(err)
}

func (h hashtagRepository) GetAll() (response []entities.Hashtag) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	cursor, err := h.Collection.Find(ctx, bson.M{})
	exceptions.PanicIfNeeded(err)

	for cursor.Next(ctx) {
		var hashtag entities.Hashtag

		err := cursor.Decode(&hashtag)
		exceptions.PanicIfNeeded(err)

		response = append(response, hashtag)
	}
	return response
}
