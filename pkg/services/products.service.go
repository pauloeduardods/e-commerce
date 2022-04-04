package services

import (
	"net/http"
	"strconv"

	"github.com/pauloeduardods/e-commerce/pkg/models"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

func GetAllProducts() ServiceResponse {
	products, err := models.GetAllProducts()
	if err != nil {
		return ServiceResponse{
			Status: http.StatusInternalServerError,
			Payload: map[string]interface{}{
				"message": "Error getting all products",
				"error":   err.Error(),
			},
		}
	}
	return ServiceResponse{
		Status: http.StatusOK,
		Payload: map[string]interface{}{
			"products": products,
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
	product, err := models.GetProduct(productID)
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
			"product": product,
		},
	}
}

func CreateProduct(product schemas.Product) ServiceResponse {
	validation := product.Validate()
	if validation.Error {
		return ServiceResponse{
			Status: validation.Status,
			Payload: map[string]interface{}{
				"message": validation.Message,
			},
		}
	}
	productID, err := models.InsertProducts(product)
	if err != nil {
		errPayload := map[string]interface{}{
			"message": "Error creating product",
			"error":   err.Error(),
		}
		return ServiceResponse{
			Status:  http.StatusInternalServerError,
			Payload: errPayload,
		}
	}
	result := schemas.Product{
		ID:       productID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}
	return ServiceResponse{
		Status:  http.StatusCreated,
		Payload: map[string]interface{}{"product": result},
	}
}
