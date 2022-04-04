package main

import (
	"github.com/go-chi/chi/v5"
)

func ProductsRoute(r chi.Router) {
	r.Get("/", GetAllProductsController)
	r.Get("/{id}", GetProductController)
	r.With(ProductValidation).Post("/", CreateProductController)
}
