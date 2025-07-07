package dto

type GetAllProductsRequest struct {
}
type GetAllProductsResponse struct {
	Id       int64
	Name     string
	Price    float64
	Discount float64
	Store    string
}
