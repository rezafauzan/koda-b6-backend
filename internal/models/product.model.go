package models

import "time"

type Product struct {
	Id          int       `json:"id"`
	Category_id int       `json:"category_id"`
	Favorite    bool      `json:"favorite_product"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Campaign_id int       `json:"campaign_id"`
	Stock       int       `json:"stock"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Images      []string  `json:"images"`
	Variants    []string  `json:"variants"`
	Portions    []string  `json:"portions"`
	Reviews     []string  `json:"reviews"`
}
