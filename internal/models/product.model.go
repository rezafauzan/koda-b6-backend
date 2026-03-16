package models

import "time"

type Product struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"category_id"`
	Favorite    bool      `json:"favorite_product"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CampaignID  int       `json:"campaign_id"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Images      []string  `json:"images"`
	Variants    []string  `json:"variants"`
	Portions    []string  `json:"portions"`
	Reviews     []string  `json:"reviews"`
}
