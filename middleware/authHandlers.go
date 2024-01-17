package middleware

import (
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"fmt"

	"glutara/repository"
	"glutara/models"
)

var (
	userRepo repository.UserRepository = repository.NewUserRepository()
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, "Hello, Welcome to Glutara Web Service!")
}

func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	existingUser, err := userRepo.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "User not found" {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		http.Error(response, "User with this email already exists", http.StatusBadRequest)
		return
	}

	userCount, err := userRepo.GetUserCount()
	if err != nil {
		http.Error(response, "Failed to get user count", http.StatusInternalServerError)
		return
	}
	user.ID = userCount + 1
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(response, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	
	_, err = userRepo.Save(&user)
	if err != nil {
		http.Error(response, "Failed to create new user", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var token models.LoginToken

	err := json.NewDecoder(request.Body).Decode(&token)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	existingUser, err := userRepo.GetUserByEmail(token.Email)
	if err != nil {
		if err.Error() == "User not found" {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		} else {
			http.Error(response, "Failed to retrieve user data", http.StatusInternalServerError)
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(token.Password))
	if err != nil {
		http.Error(response, "Email and password do not match", http.StatusUnauthorized)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(existingUser)
}