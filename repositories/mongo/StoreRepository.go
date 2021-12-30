package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-clean-arch/configurations"
	"golang-clean-arch/entities"
	"golang-clean-arch/exceptions"
	"golang-clean-arch/usecases/out"
)

type storeRepository struct {
	Collection *mongo.Collection
}

func NewStoreRepository(database *mongo.Database) out.StoreDataSource {
	return &storeRepository{
		Collection: database.Collection("stores"),
	}
}

func (s storeRepository) Create(store entities.Store) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	_, err := s.Collection.InsertOne(ctx, bson.M{
		"_id":      store.Id,
		"name":     store.Name,
		"image":    store.Image,
		"location": store.Location,
		"lat":      store.Lat,
		"long":     store.Long,
	})

	exceptions.PanicIfNeeded(err)
}

func (s storeRepository) GetStore(id string) (response entities.Store) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	err := s.Collection.FindOne(
		ctx,
		bson.D{
			{"_id", id},
		},
	).Decode(&response)

	exceptions.PanicIfNeeded(err)
	return response
}
