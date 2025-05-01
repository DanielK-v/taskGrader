package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	models "github.com/DanielK_v/taskGrader/models/users"
)

func GetAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to fetch users"},
		)
		return
	}

	context.JSON(http.StatusOK, users)
}

func Register(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Please, try again!"})
		return
	}
	user.Password = string(hashedPassword)

	_, err = models.AddUser(user)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			// Duplicate entry
			context.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
