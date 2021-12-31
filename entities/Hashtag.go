package entities

type Hashtag struct {
	Id    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Count int8   `json:"count" bson:"count"`
}
