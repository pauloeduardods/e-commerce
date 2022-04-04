package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/pauloeduardods/e-commerce/pkg/schemas"
	"github.com/pauloeduardods/e-commerce/pkg/services"
)

func GetAllProductsController(w http.ResponseWriter, r *http.Request) {
	serviceResponse := services.GetAllProducts()
	render.Status(r, serviceResponse.Status)
	render.JSON(w, r, serviceResponse.Payload)
}

func GetProductController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	serviceResponse := services.GetProduct(id)
	render.Status(r, serviceResponse.Status)
	render.JSON(w, r, serviceResponse.Payload)
}

func CreateProductController(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value(schemas.Product{})
	if p == nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "Invalid request"})
		return
	}
	product := p.(schemas.Product)
	serviceResponse := services.CreateProduct(product)
	render.Status(r, serviceResponse.Status)
	render.JSON(w, r, serviceResponse.Payload)
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(schemas.User{})
	if u == nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "Invalid request"})
		return
	}
	user := u.(schemas.User)
	serviceResponse := services.Login(user.Email, user.Password, AppConfig.HmacSecret)
	render.Status(r, serviceResponse.Status)
	render.JSON(w, r, serviceResponse.Payload)
}
