package service

import (
	"go-product-app/domain"
	"go-product-app/service"
	"go-product-app/service/request"
	"os"
	"testing"
)

var productService service.ProductService

func TestMain(m *testing.M) {

	initialProducts := []domain.Product{
		{1, "AirFryer", 1000.0, 0.0, "A_TECH"},
		{2, "Oven", 1999.0, 0.0, "B_TECH"},
	}

	fakeProductRepository := newFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)

	exitCode := m.Run()
	os.Exit(exitCode)

}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		actualProducts := productService.GetAll(request.GetAllProductsRequest{})
		if len(actualProducts.Items) != 2 {
			t.Errorf("expected 2 products, got %d", len(actualProducts.Items))
		}
		if actualProducts.Items[0].Name != "AirFryer" || actualProducts.Items[1].Name != "Oven" {
			t.Errorf("unexpected product names: %+v", actualProducts.Items)
		}
	})
}

func Test_ShouldAddProduct(t *testing.T) {
	t.Run("ShouldAddProduct", func(t *testing.T) {
		addReq := request.AddProductRequest{
			Name:     "Toaster",
			Price:    500.0,
			Discount: 10.0,
			Store:    "C_TECH",
		}
		addedProduct, err := productService.Add(addReq)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if addedProduct.Name != addReq.Name || addedProduct.Price != addReq.Price || addedProduct.Store != addReq.Store {
			t.Errorf("added product does not match request: %+v", addedProduct)
		}
		// Check if product count increased
		allProducts := productService.GetAll(request.GetAllProductsRequest{})
		if len(allProducts.Items) != 3 {
			t.Errorf("expected 3 products after add, got %d", len(allProducts.Items))
		}
	})

	t.Run("ShouldReturnErrorWhenDiscountExceedsSeventy", func(t *testing.T) {
		addReq := request.AddProductRequest{
			Name:     "Microwave",
			Price:    800.0,
			Discount: 75.0,
			Store:    "D_TECH",
		}
		_, err := productService.Add(addReq)
		if err == nil {
			t.Error("expected error when discount exceeds 70, got nil")
		}
	})
}
