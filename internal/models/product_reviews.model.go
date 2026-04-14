package models

import "time"

type ProductReview struct {
	Id        int       `json:"id" db:"id"`
	ProductId int       `json:"product_id" db:"product_id"`
	UserId    int       `json:"user_id" db:"user_id"`
	Rating    int       `json:"rating" db:"rating"`
	Messages  string    `json:"messages" db:"messages"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
