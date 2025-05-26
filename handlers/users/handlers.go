package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"github.com/DanielK_v/taskGrader/models"
	"github.com/DanielK_v/taskGrader/utils"
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

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Please, try again!"})
		return
	}
	user.Password = string(bytes)

	_, err = models.AddUser(user)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(context *gin.Context) {
	var loginRequest models.LoginRequest

	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	user, err := models.GetUserByEmail(loginRequest.Email)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the user"})
	}

	err = utils.CheckPasswordHash(user, loginRequest.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}
		
	token, err := utils.GenerateToken(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed generating the access token"})
		return
	}
			

	context.JSON(http.StatusOK, gin.H{"access_token": token})
}
