package models

import "time"

type CartItem struct {
	Id        int        `json:"id" db:"id"`
	CartId    int        `json:"cart_id" db:"cart_id"`
	ProductId int        `json:"product_id" db:"product_id"`
	Size      string     `json:"size" db:"size"`
	Hotice    string     `json:"hotice" db:"hotice"`
	Quantity  int        `json:"quantity" db:"quantity"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}