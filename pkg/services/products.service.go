package services

import (
	"net/http"
	"strconv"

	"github.com/pauloeduardods/e-commerce/pkg/models"
)

func GetAllProducts() ServiceResponse {
	// fmt.Println(models.GetAllProducts())
	// fmt.Println(models.InsertProducts(models.Product{Name: "testeeeeee", Quantity: 1, Price: 1.0}))
	// fmt.Println(models.GetAllProducts())
	// fmt.Println(models.GetProduct(2))
	products, err := models.GetAllProducts()
	if err != nil {
		errPayload := map[string]interface{}{
			"message": "Error getting all products",
			"error":   err.Error(),
		}
		return ServiceResponse{Status: http.StatusInternalServerError, Payload: errPayload}
	}
	return ServiceResponse{Status: http.StatusOK, Payload: products}
}

func GetProduct(id string) ServiceResponse {
	productID, err := strconv.Atoi(id)
	if err != nil {
		errPayload := map[string]interface{}{
			"message": "Error converting product id",
			"error":   err.Error(),
		}
		return ServiceResponse{Status: http.StatusBadRequest, Payload: errPayload}
	}
	product, err := models.GetProduct(productID)
	if err != nil {
		emptyPayload := map[string]interface{}{}
		return ServiceResponse{Status: http.StatusNotFound, Payload: emptyPayload}
	}
	return ServiceResponse{Status: http.StatusOK, Payload: product}
}
