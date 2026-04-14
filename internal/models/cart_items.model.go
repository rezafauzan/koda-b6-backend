package models

import "time"

type CartItem struct {
	Id        int        `json:"id" db:"id"`
	CartId    int        `json:"cart_id" db:"cart_id"`
	ProductId int        `json:"product_id" db:"product_id"`
	SizeId    int        `json:"size_id" db:"size_id"`
	VariantId int        `json:"variant_id" db:"variant_id"`
	Quantity  int        `json:"quantity" db:"quantity"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}