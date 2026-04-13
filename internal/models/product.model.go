package models

import "time"

type Product struct {
	Id              int       `json:"id" db:"id"`
	CategoryId      int       `json:"category_id" db:"category_id"`
	Name            string    `json:"name" db:"name"`
	Description     string    `json:"description" db:"description"`
	Price           float64       `json:"price" db:"price"`
	Stock           int       `json:"stock" db:"stock"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
