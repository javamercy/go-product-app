package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_PRODUCTS = `INSERT INTO products (name, price, discount, store)
	VALUES
	    ('Oven', 1000, 10, 'A TECH'),
		('Refrigerator', 2000, 20, 'A TECH'),
		('Washing Machine', 1500, 15, 'B TECH'),
		('Microwave', 800, 5, 'B TECH')`

func TestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertProductResult, insertProductErr := dbPool.Exec(ctx, INSERT_PRODUCTS)
	if insertProductErr != nil {
		log.Error(insertProductErr)
	} else {
		log.Info(fmt.Sprintf("products data created with %d rows", insertProductResult.RowsAffected()))
	}
}
