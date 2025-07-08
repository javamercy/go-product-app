package request

type AddProductRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
	Store    string  `json:"store"`
}

type GetAllProductsRequest struct {
}
