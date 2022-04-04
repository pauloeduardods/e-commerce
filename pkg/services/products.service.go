package services

import (
	"net/http"
	"strconv"

	"github.com/pauloeduardods/e-commerce/pkg/models"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

func GetAllProducts() ServiceResponse {
	products := make(chan []schemas.Product)
	go models.GetAllProducts(products)
	return ServiceResponse{
		Status: http.StatusOK,
		Payload: map[string]interface{}{
			"products": <-products,
		},
	}
}

func GetProduct(id string) ServiceResponse {
	productID, err := strconv.Atoi(id)
	if err != nil {
		return ServiceResponse{
			Status: http.StatusBadRequest,
			Payload: map[string]interface{}{
				"message": "Error converting product id",
				"error":   err.Error(),
			},
		}
	}
	product := make(chan schemas.Product)
	go models.GetProduct(productID, product)
	if err != nil {
		emptyPayload := map[string]interface{}{}
		return ServiceResponse{
			Status:  http.StatusNotFound,
			Payload: emptyPayload,
		}
	}
	return ServiceResponse{
		Status: http.StatusOK,
		Payload: map[string]interface{}{
			"product": <-product,
		},
	}
}

func CreateProduct(product schemas.Product) ServiceResponse {
	productID := make(chan int64)
	go models.InsertProducts(product, productID)
	id := <-productID
	if id == 0 {
		return ServiceResponse{
			Status: http.StatusInternalServerError,
			Payload: map[string]interface{}{
				"message": "Error creating product",
			},
		}
	}
	result := schemas.Product{
		ID:       id,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}
	return ServiceResponse{
		Status:  http.StatusCreated,
		Payload: map[string]interface{}{"product": result},
	}
}
