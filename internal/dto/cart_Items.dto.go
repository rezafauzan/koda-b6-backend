package dto
type CreateCartItemRequestDTO struct {
	CartId    int    `json:"cart_id"`
	ProductId int    `json:"product_id"`
	SizeId    int    `json:"size_id"`
	VariantId int    `json:"variant_id"`
	Quantity  int    `json:"quantity"`
}
