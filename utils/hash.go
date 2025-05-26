package utils

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/DanielK_v/taskGrader/models"
)


func CheckPasswordHash(user *models.User, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
}
