package services

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pauloeduardods/e-commerce/pkg/models"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

func GetAllProducts() ServiceResponse {
	productsChan := make(chan []schemas.Product)
	go models.GetAllProducts(productsChan)
	var products []schemas.Product
	select {
	case products = <-productsChan:
		return ServiceResponse{
			Status: http.StatusOK,
			Payload: map[string]interface{}{
				"products": products,
			},
		}
	case <-time.After(time.Second * 5):
		return ServiceResponse{
			Status: http.StatusGatewayTimeout,
			Payload: map[string]interface{}{
				"message": "Error getting products",
			},
		}
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
	productChan := make(chan schemas.Product)
	go models.GetProduct(productID, productChan)
	var product schemas.Product
	select {
	case product = <-productChan:
		return ServiceResponse{
			Status: http.StatusOK,
			Payload: map[string]interface{}{
				"product": product,
			},
		}
	case <-time.After(time.Second * 5):
		return ServiceResponse{
			Status: http.StatusGatewayTimeout,
			Payload: map[string]interface{}{
				"message": "Error getting product",
			},
		}
	}
}

func CreateProduct(product schemas.Product) ServiceResponse {
	productID := make(chan int64)
	go models.InsertProduct(product, productID)
	select {
	case id := <-productID:
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
	case <-time.After(time.Second * 5):
		return ServiceResponse{
			Status: http.StatusGatewayTimeout,
			Payload: map[string]interface{}{
				"message": "Error creating product",
			},
		}
	}
}
