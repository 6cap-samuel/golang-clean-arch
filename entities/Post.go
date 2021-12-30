package entities

type Post struct {
	Id          string `json:"id" bson:"_id"`
	Description string `json:"description" bson:"description"`
	Store       Store  `json:"store" bson:"store"`
	Rating      int8   `json:"rating" bson:"rating"`
}
