package item

import "time"

type Item struct {
	ID         string     `json:"id"`
	Price      int64      `json:"price"`
	Name       string     `json:"name"`
	PictureURL string     `json:"pictureUrl"`
	Stock      int64      `json:"stock"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	Featured   bool       `json:"featured"`
	Visits     int        `json:"visits"`
}

func (i *Item) GetID() string {
	if i != nil {
		return i.ID
	}
	return ""
}

func (i *Item) GetPrice() int64 {
	if i != nil {
		return i.Price
	}
	return 0
}

func (i *Item) GetName() string {
	if i != nil {
		return i.Name
	}
	return ""
}

func (i *Item) GetPicture() string {
	if i != nil {
		return i.PictureURL
	}
	return ""
}

func (i *Item) GetStock() int64 {
	if i != nil {
		return i.Stock
	}
	return 0
}
