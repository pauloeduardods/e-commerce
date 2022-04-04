package services

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pauloeduardods/e-commerce/pkg/models"
	"github.com/pauloeduardods/e-commerce/pkg/schemas"
)

func Login(email, password string, HmacSecret []byte) ServiceResponse {
	userChan := make(chan schemas.User)
	models.GetUserByEmail(email, userChan)
	user := <-userChan
	if user.Password != password {
		return ServiceResponse{
			Status: http.StatusUnauthorized,
			Payload: map[string]interface{}{
				"message": "Invalid credentials",
			},
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(HmacSecret)
	if err != nil {
		return ServiceResponse{
			Status: http.StatusInternalServerError,
			Payload: map[string]interface{}{
				"message": "Error generating token",
			},
		}
	}
	return ServiceResponse{
		Status: http.StatusOK,
		Payload: map[string]interface{}{
			"token": tokenString,
		},
	}
}
