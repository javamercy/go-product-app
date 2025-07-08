package service

import (
	"errors"
	commonResponse "go-product-app/common/response"
	"go-product-app/domain"
	"go-product-app/persistence"
	"go-product-app/service/helper"
	serviceRequest "go-product-app/service/request"
	serviceResponse "go-product-app/service/response"
)

type ProductService interface {
	GetAll(request serviceRequest.GetAllProductsRequest) commonResponse.ListResponse[serviceResponse.GetAllProductsResponse]
	GetAllByStore(storeName string) []domain.Product
	Add(request serviceRequest.AddProductRequest) (serviceResponse.AddedProductResponse, error)
	GetById(productId int64) (domain.Product, error)
	Update(product domain.Product) error
	DeleteById(productId int64) error
}

func NewProductService(productRepository persistence.ProductRepository) ProductService {
	return &ProductManager{productRepository: productRepository}
}

type ProductManager struct {
	productRepository persistence.ProductRepository
}

func (productManager *ProductManager) GetAll(request serviceRequest.GetAllProductsRequest) commonResponse.ListResponse[serviceResponse.GetAllProductsResponse] {
	products := productManager.productRepository.GetAll()

	var getAllProductsResponse commonResponse.ListResponse[serviceResponse.GetAllProductsResponse]

	for _, product := range products {

		getAllProductsResponse.Items = append(getAllProductsResponse.Items, serviceResponse.GetAllProductsResponse{
			Id:       product.Id,
			Name:     product.Name,
			Store:    product.Store,
			Discount: product.Discount,
			Price:    product.Price,
		})
	}

	return getAllProductsResponse

}

func (productManager *ProductManager) GetAllByStore(storeName string) []domain.Product {
	//TODO implement me
	panic("implement me")
}

func (productManager *ProductManager) Add(request serviceRequest.AddProductRequest) (serviceResponse.AddedProductResponse, error) {

	err := discountCannotExceedSeventy(request.Discount)

	if err != nil {
		return serviceResponse.AddedProductResponse{}, err
	}

	product := helper.ToProduct(request)

	response, err := productManager.productRepository.Add(product)

	if err != nil {
		return serviceResponse.AddedProductResponse{}, err
	}

	return helper.ToAddedProductResponse(response), nil
}

func (productManager *ProductManager) GetById(productId int64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (productManager *ProductManager) Update(product domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func (productManager *ProductManager) DeleteById(productId int64) error {
	//TODO implement me
	panic("implement me")
}

func discountCannotExceedSeventy(discount float64) error {
	if discount > 70.0 {
		return errors.New("discount cannot exceed 70.0")
	}
	return nil
}
