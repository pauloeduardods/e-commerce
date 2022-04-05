package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

////////////////////////////////////////////////////////////////////////////////////////////
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"message": "Unauthorized",
			})
			return
		}
		// ctx := context.WithValue(r.Context(), "token", tokenString)
		// next.ServeHTTP(w, r.WithContext(ctx))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return AppConfig.HmacSecret, nil
		})
		fmt.Println(11)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
			fmt.Println(claims)
		} else {
			println(12)
			fmt.Println(err)
		}
	})
}

////////////////////////////////////////////////////////////////////////////////////////////

func ProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product schemas.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}
		validation := product.Validate()
		if validation.Error {
			render.Status(r, validation.Status)
			render.JSON(w, r, map[string]string{
				"message": validation.Message,
			})
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
			render.JSON(w, r, map[string]string{
				"message": "Invalid request",
				"error":   err.Error(),
			})
			return
		}
		validation := login.Validate()
		if validation.Error {
			render.Status(r, validation.Status)
			render.JSON(w, r, map[string]string{
				"message": validation.Message,
			})
			return
		}
		ctx := context.WithValue(r.Context(), schemas.User{}, login)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
