package dto

import "time"

type ProductResponseDTO struct {
	Id              int       `json:"id"`
	CategoryId      int       `json:"category_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           int       `json:"price"`
	Stock           int       `json:"stock"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}