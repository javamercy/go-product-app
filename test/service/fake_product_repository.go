package service

import (
	"go-product-app/domain"
	"go-product-app/persistence"
)

type FakeProductRepository struct {
	Products []domain.Product
}

func newFakeProductRepository(initialProducts []domain.Product) persistence.ProductRepository {
	return &FakeProductRepository{
		Products: initialProducts,
	}
}

func (productRepository *FakeProductRepository) GetAll() []domain.Product {

	return productRepository.Products
}
func (productRepository *FakeProductRepository) GetAllByStore(storeName string) []domain.Product {
	//TODO implement me
	panic("implement me")
}

func (productRepository *FakeProductRepository) Add(product domain.Product) (domain.Product, error) {

	newProduct := domain.Product{
		Id:       int64(len(productRepository.Products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Store:    product.Store,
		Discount: product.Discount,
	}
	productRepository.Products = append(productRepository.Products, newProduct)

	return newProduct, nil
}

func (productRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (productRepository *FakeProductRepository) Update(product domain.Product) error {
	//TODO implement me
	panic("implement me")
}

func (productRepository *FakeProductRepository) DeleteById(productId int64) error {
	//TODO implement me
	panic("implement me")
}
