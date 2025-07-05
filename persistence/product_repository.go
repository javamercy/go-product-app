package persistence

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"go-product-app/domain"
)

type ProductRepository interface {
	GetAll() []domain.Product
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

	var products []domain.Product
	var id int64
	var name string
	var price float64
	var discount float64
	var store string

	for rows.Next() {
		err := rows.Scan(&id, &name, &price, &discount, &store)

		if err != nil {
			log.Error("Error while scanning products: %v", err)
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
