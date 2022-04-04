package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"message": "Unauthorized"})
			return
		}
		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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

func LoginValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var login schemas.User
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"message": "Invalid request", "error": err.Error()})
			return
		}
		validation := login.Validate()
		if validation.Error {
			render.Status(r, validation.Status)
			render.JSON(w, r, map[string]string{"message": validation.Message})
			return
		}
		ctx := context.WithValue(r.Context(), schemas.User{}, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
