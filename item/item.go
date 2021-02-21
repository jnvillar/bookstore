package item

type Item struct {
	ID         string `json:"id"`
	Price      int64  `json:"price"`
	Name       string `json:"name"`
	PictureURL string `json:"picture_url"`
	Stock      int64  `json:"stock"`
}
