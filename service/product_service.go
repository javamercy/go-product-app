package service

import (
	"errors"
	"go-product-app/common/response"
	"go-product-app/domain"
	"go-product-app/persistence"
	"go-product-app/service/dto"
)

type ProductService interface {
	GetAll(request dto.GetAllProductsRequest) response.ListResponse[dto.GetAllProductsResponse]
	GetAllByStore(storeName string) []domain.Product
	Add(request dto.AddProductRequest) (dto.AddedProductResponse, error)
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

func (productManager *ProductManager) GetAll(request dto.GetAllProductsRequest) response.ListResponse[dto.GetAllProductsResponse] {
	products := productManager.productRepository.GetAll()

	var getAllProductsResponse response.ListResponse[dto.GetAllProductsResponse]

	for _, product := range products {

		getAllProductsResponse.Items = append(getAllProductsResponse.Items, dto.GetAllProductsResponse{
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

func (productManager *ProductManager) Add(request dto.AddProductRequest) (dto.AddedProductResponse, error) {

	err := discountCannotExceedSeventy(request.Discount)

	if err != nil {
		return dto.AddedProductResponse{}, err
	}

	product := domain.Product{
		Name:     request.Name,
		Price:    request.Price,
		Discount: request.Discount,
		Store:    request.Store,
	}

	response, err := productManager.productRepository.Add(product)

	if err != nil {
		return dto.AddedProductResponse{}, err
	}

	return dto.AddedProductResponse{
		Id:       response.Id,
		Name:     response.Name,
		Price:    response.Price,
		Discount: response.Discount,
		Store:    request.Store,
	}, nil
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
