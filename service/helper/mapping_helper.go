package helper

import (
	"go-product-app/domain"
	serviceRequest "go-product-app/service/request"
	serviceResponse "go-product-app/service/response"
)

func ToProduct(req serviceRequest.AddProductRequest) domain.Product {
	return domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Discount: req.Discount,
		Store:    req.Store,
	}
}

func ToAddedProductResponse(product domain.Product) serviceResponse.AddedProductResponse {
	return serviceResponse.AddedProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}
