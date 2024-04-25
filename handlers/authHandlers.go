package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"glutara/models"
	"glutara/repository"
)

func MainHandler(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello, Welcome to Glutara Web Service!")
}

func Register(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	existingUser, err := repository.UserRepo.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "user not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to register user"})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Email already used"})
		return
	}

	userCount, err := repository.UserRepo.GetUserCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to register user"})
		return
	}
	user.ID = userCount + 1
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to register user"})
		return
	}
	user.Password = string(hashedPassword)

	_, err = repository.UserRepo.Save(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	existingUser, err := repository.UserRepo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to log user"})
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Wrong user credentials"})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}
