package mongo

import (
	"feature-service/configurations"
	"feature-service/entities"
	"feature-service/exceptions"
	"feature-service/usecases/out"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type featureRepository struct {
	Collection *mongo.Collection
}

func NewFeatureRepository(database *mongo.Database) out.FeatureDataSource {
	return &featureRepository{
		Collection: database.Collection("features"),
	}
}

func (f featureRepository) GetAll() (response []entities.Feature) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	findOptions := options.Find()

	cursor, err := f.Collection.Find(ctx, bson.M{}, findOptions)
	exceptions.PanicIfNeeded(err)

	for cursor.Next(ctx) {
		var feature entities.Feature

		err := cursor.Decode(&feature)
		exceptions.PanicIfNeeded(err)

		response = append(response, feature)
	}
	return response
}
