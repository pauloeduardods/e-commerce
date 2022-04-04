package services

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pauloeduardods/e-commerce/pkg/models"
)

func Login(email, password string, HmacSecret []byte) ServiceResponse {
	user, err := models.GetUserByEmail(email)
	if err != nil || user.Password != password {
		return ServiceResponse{Status: http.StatusUnauthorized, Payload: map[string]interface{}{"message": "Invalid credentials"}}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"time": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(HmacSecret)
	if err != nil {
		return ServiceResponse{Status: http.StatusInternalServerError, Payload: map[string]interface{}{"message": "Error generating token"}}
	}
	return ServiceResponse{Status: http.StatusOK, Payload: map[string]interface{}{"token": tokenString}}
}
