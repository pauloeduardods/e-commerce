package main

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/pauloeduardods/e-commerce/pkg/services"
)

func GetAllProductsController(w http.ResponseWriter, r *http.Request) {
	serviceResponse := services.GetAllProducts()
	render.Status(r, serviceResponse.Status)
	render.JSON(w, r, serviceResponse.Payload)
}
