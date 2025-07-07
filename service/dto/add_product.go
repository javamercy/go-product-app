package dto

type AddProductRequest struct {
	Name     string
	Price    float64
	Discount float64
	Store    string
}

type AddedProductResponse struct {
	Id       int64
	Name     string
	Price    float64
	Discount float64
	Store    string
}
