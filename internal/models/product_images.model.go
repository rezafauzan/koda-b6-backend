package models

import "time"

type ProductImages struct {
	Id         int       `json:"id"`
	Product_id int       `json:"product_id"`
	Image      string    `json:"image"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}