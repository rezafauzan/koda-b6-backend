package dto

type CreatePaymentRequestDTO struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Delivery string `json:"delivery"`
}
