package utils

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/DanielK_v/taskGrader/models"
	

)

var secretKey string = os.Getenv("KEY")

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"userId": user.Id,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}