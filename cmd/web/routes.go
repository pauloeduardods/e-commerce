package main

import "github.com/go-chi/chi/v5"

func ProductsRoute(r chi.Router) {
	r.Get("/", GetAllProductsController)
	r.Get("/{id}", GetProductController)
	r.With(AuthMiddleware).With(ProductValidation).Post("/", CreateProductController)
}

func LoginRoute(r chi.Router) {
	r.With(LoginValidation).Post("/", LoginController)
}
