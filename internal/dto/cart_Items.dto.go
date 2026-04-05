package dto
type CreateCartItemRequestDTO struct {
	CartId    int    `json:"cart_id"`
	ProductId int    `json:"product_id"`
	Size      string `json:"size"`
	Hotice    string `json:"hotice"`
	Quantity  int    `json:"quantity"`
}
