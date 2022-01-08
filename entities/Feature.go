package entities

type Feature struct {
	Id          string   `bson:"_id" json:"_id"`
	Name        string   `bson:"name" json:"name"`
	IsFeatureOn bool     `bson:"is_feature_on" json:"is_feature_on"`
	Roles       []string `bson:"roles" json:"roles"`
}
