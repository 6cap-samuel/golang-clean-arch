package entities

type Store struct {
	Id       string  `bson:"_id"`
	Name     string  `bson:"name"`
	Image    string  `bson:"image"`
	Location string  `bson:"location"`
	Lat      float32 `bson:"lat"`
	Long     float32 `bson:"long"`
}
