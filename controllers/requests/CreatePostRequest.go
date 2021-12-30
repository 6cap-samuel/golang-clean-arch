package requests

type CreatePostRequest struct {
	StoreName     string  `json:"store_name"`
	StoreImage    string  `json:"store_image"`
	StoreLocation string  `json:"store_location"`
	StoreLat      float32 `json:"store_lat"`
	StoreLong     float32 `json:"store_long"`
	Description   string  `json:"description"`
}
