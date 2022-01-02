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

// TODO: refactor this piece of code
func (h hashtagRepository) Create(hashtags []entities.Hashtag) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	// get all hashtags by keywords
	var conditions []string
	for _, hashtag := range hashtags {
		conditions = append(conditions, hashtag.Name)
	}
	var response = h.GetAllByName(conditions)

	if len(response) != 0 {
		// filter hashtag: O(N2) to find out better algorithm
		var filteredHashtags []entities.Hashtag
		for _, hashtag := range response {
			var isFound = false
			for _, cur := range hashtags {
				if hashtag.Name == cur.Name {
					isFound = true
					break
				}
			}
			if !isFound {
				filteredHashtags = append(filteredHashtags, hashtag)
			}
		}

		if len(filteredHashtags) != 0 {
			// insert filtered hashtags
			var insertionInterface []interface{}
			for _, t := range filteredHashtags {
				insertionInterface = append(insertionInterface, t)
			}
			_, err := h.Collection.InsertMany(ctx, insertionInterface)
			exceptions.PanicIfNeeded(err)
		}
	} else {
		var insertionInterface []interface{}
		for _, t := range hashtags {
			insertionInterface = append(insertionInterface, t)
		}
		_, err := h.Collection.InsertMany(ctx, insertionInterface)
		exceptions.PanicIfNeeded(err)
	}

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

func (h hashtagRepository) GetAllByName(vals []string) (response []entities.Hashtag) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	var bsonMap bson.A
	for _, item := range vals {
		bsonMap = append(bsonMap, item)
	}

	cursor, err := h.Collection.Find(
		ctx,
		bson.D{
			{"name",
				bson.D{
					{"$in",
						bsonMap,
					},
				},
			},
		},
	)

	exceptions.PanicIfNeeded(err)

	for cursor.Next(ctx) {
		var hashtag entities.Hashtag

		err := cursor.Decode(&hashtag)
		exceptions.PanicIfNeeded(err)

		response = append(response, hashtag)
	}
	return response
}
