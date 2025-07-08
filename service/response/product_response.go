package response

type AddedProductResponse struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
	Store    string  `json:"store"`
}

type GetAllProductsResponse struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
	Store    string  `json:"store"`
}
