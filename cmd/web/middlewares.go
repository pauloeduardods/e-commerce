package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

func ProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product schemas.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"message": "Invalid request", "error": err.Error()})
			return
		}
		validation := product.Validate()
		if validation.Error {
			render.Status(r, validation.Status)
			render.JSON(w, r, map[string]string{"message": validation.Message})
			return
		}
		ctx := context.WithValue(r.Context(), schemas.Product{}, product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
