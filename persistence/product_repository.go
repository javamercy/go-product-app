package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"go-product-app/domain"
	"go-product-app/persistence/common"
)

type ProductRepository interface {
	GetAll() []domain.Product
	GetAllByStore(storeName string) []domain.Product
	Add(product domain.Product) (domain.Product, error)
	GetById(productId int64) (domain.Product, error)
	Update(product domain.Product) error
	DeleteById(productId int64) error
}

type PostgresProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) ProductRepository {
	return &PostgresProductRepository{dbPool: dbPool}
}

func (productRepository *PostgresProductRepository) GetAll() []domain.Product {
	ctx := context.Background()
	rows, err := productRepository.dbPool.Query(ctx, "select * from products")

	if err != nil {
		log.Error("Error while fetching products: %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(rows)
}

func (productRepository *PostgresProductRepository) GetAllByStore(storeName string) []domain.Product {
	sql := `select * from products where Store = $1`

	rows, err := productRepository.dbPool.Query(context.Background(), sql, storeName)

	if err != nil {
		log.Error("error while running GetAllByStore:", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(rows)

}

func (productRepository *PostgresProductRepository) Add(product domain.Product) (domain.Product, error) {
	sql := `insert into products (name, price, discount, store) values ($1, $2, $3, $4)`

	var id int64
	err := productRepository.dbPool.QueryRow(context.Background(), sql,
		product.Name,
		product.Price,
		product.Discount,
		product.Store).Scan(&id)

	if err != nil {
		log.Error(err)
		return domain.Product{}, err
	}

	return domain.Product{
		Id:       id,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}, nil
}

func (productRepository *PostgresProductRepository) GetById(productId int64) (domain.Product, error) {

	sql := `select * from products where id = $1`
	row := productRepository.dbPool.QueryRow(context.Background(), sql, productId)

	var id int64
	var name string
	var price float64
	var discount float64
	var store string

	err := row.Scan(&id, &name, &price, &discount, &store)

	if err != nil && err.Error() == common.NOT_FOUND {
		log.Error(err)
		return domain.Product{}, errors.New(fmt.Sprintf("Product with id %v not found", productId))
	}

	product := domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}

	return product, nil

}

func (productRepository *PostgresProductRepository) Update(product domain.Product) error {

	sql := `update products set name = $1, price = $2, discount = $3, store = $4 where id = $5`

	_, err := productRepository.dbPool.Exec(context.Background(), sql,
		product.Name,
		product.Price,
		product.Discount,
		product.Store,
		product.Id)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (productRepository *PostgresProductRepository) DeleteById(productId int64) error {
	sql := `delete from products where id = $1`

	_, err := productRepository.dbPool.Exec(context.Background(), sql, productId)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func extractProductsFromRows(rows pgx.Rows) []domain.Product {
	var products []domain.Product
	var id int64
	var name string
	var price float64
	var discount float64
	var store string

	for rows.Next() {
		err := rows.Scan(&id, &name, &price, &discount, &store)

		if err != nil {
			log.Error(err)
		} else {
			products = append(products, domain.Product{
				Id:       id,
				Name:     name,
				Price:    price,
				Discount: discount,
				Store:    store,
			})
		}
	}
	return products
}
