package controller

import (
	"github.com/labstack/echo/v4"
	"go-product-app/service"
	"go-product-app/service/request"
	"net/http"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {

	return &ProductController{
		productService: productService,
	}
}

func (productController ProductController) MapRoutes(e *echo.Echo) {

	e.GET("/api/v1/products", productController.GetAll)
	e.GET("/api/v1/products/:id", productController.GetById)
	e.POST("/api/v1/products", productController.Add)
	e.PUT("/api/v1/products", productController.Update)
	e.DELETE("/api/v1/products", productController.Delete)
}

func (productController ProductController) GetAll(c echo.Context) error {
	store := c.QueryParam("store")
	if len(store) == 0 {
		products := productController.productService.GetAll(request.GetAllProductsRequest{})
		return c.JSON(http.StatusOK, products)
	} else {
		return nil
	}

}

func (productController ProductController) GetById(c echo.Context) error {
	// TODO: implement me
	panic("implement me")

}

func (productController ProductController) Add(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	_ = c.Bind(&addProductRequest)

	addedProductResponse, err := productController.productService.Add(addProductRequest)

	if err == nil {
		return c.JSON(http.StatusCreated, addedProductResponse)
	}
	return nil

}

func (productController ProductController) Update(c echo.Context) error {
	// TODO: implement me
	panic("implement me")

}

func (productController ProductController) Delete(c echo.Context) error {
	// TODO: implement me
	panic("implement me")
}
