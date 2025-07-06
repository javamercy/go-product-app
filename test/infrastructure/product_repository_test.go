package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"go-product-app/common/postgresql"
	"go-product-app/domain"
	"go-product-app/persistence"
	"os"
	"testing"
)

var productRepository persistence.ProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		UserName:              "postgres",
		Password:              "postgres",
		DatabaseName:          "productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("Before all tests")
	code := m.Run()
	fmt.Println("After all tests")
	os.Exit(code)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {

	expectedProducts := []domain.Product{
		{Id: 1, Name: "Oven", Price: 1000.0, Discount: 10.0, Store: "A_TECH"},
		{Id: 2, Name: "Refrigerator", Price: 2000.0, Discount: 20.0, Store: "A_TECH"},
		{Id: 3, Name: "Washing Machine", Price: 1500.0, Discount: 15.0, Store: "B_TECH"},
		{Id: 4, Name: "Microwave", Price: 800.0, Discount: 5.0, Store: "B_TECH"},
	}
	setup(ctx, dbPool)
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAll()
		assert.Equal(t, len(expectedProducts), len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	expectedProducts := []domain.Product{
		{Id: 3, Name: "Washing Machine", Price: 1500.0, Discount: 15.0, Store: "B_TECH"},
		{Id: 4, Name: "Microwave", Price: 800.0, Discount: 5.0, Store: "B_TECH"},
	}
	setup(ctx, dbPool)
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllByStore("B_TECH")
		assert.Equal(t, len(expectedProducts), len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})

	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	product := domain.Product{
		Name: "Washing Machine", Price: 1500.0, Discount: 15.0, Store: "B_TECH",
	}
	expectedProducts :=
		[]domain.Product{
			{Id: 1, Name: "Washing Machine", Price: 1500.0, Discount: 15.0, Store: "B_TECH"},
		}

	t.Run("AddProduct", func(t *testing.T) {

		productRepository.Add(product)
		actualProducts := productRepository.GetAll()
		assert.Equal(t, len(expectedProducts), len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)

	})

	clear(ctx, dbPool)
}

func TestGetProductById(t *testing.T) {
	setup(ctx, dbPool)

	expectedProduct := domain.Product{
		Id: 1, Name: "Oven", Price: 1000.0, Discount: 10.0, Store: "A_TECH",
	}

	expectedErr := "Product with id 5 not found"

	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetById(1)
		_, actualErr := productRepository.GetById(5)

		assert.Equal(t, expectedProduct, actualProduct)
		assert.Equal(t, expectedErr, actualErr.Error())
	})

	clear(ctx, dbPool)
}
