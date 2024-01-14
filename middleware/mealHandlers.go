package middleware

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

	"glutara/repository"
	"glutara/models"
)

var (
	mealRepo repository.MealRepository = repository.NewMealRepository()
)

func GetMeals(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	meals, err := mealRepo.FindAllUserMeals(userID)
	if err != nil {
		http.Error(response, "Failed to retrieve meal logs data", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(meals)
}

func CreateMeal(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var meal models.Meal
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&meal)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != meal.UserID {
		http.Error(response, "Request payload does not match", http.StatusBadRequest)
		return
	}

	_, err = userRepo.GetUserByID(userID)
	if err != nil {
		if err.Error() != "User not found" {
			http.Error(response, "Failed to check user", http.StatusInternalServerError)
			return
		} else {
			http.Error(response, "User does not exist", http.StatusBadRequest)
			return
		}
	}

	max, err := mealRepo.GetUserMealsMaxCount(userID)
	if err != nil {
		http.Error(response, "Failed to get max meal log count", http.StatusInternalServerError)
		return
	}
	meal.MealID = max + 1

	_, err = mealRepo.Save(&meal)
	if err != nil {
		http.Error(response, "Failed to create new meal log", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(meal)
}

func DeleteMeal(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	mealID, err := strconv.ParseInt(vars["MealID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid meal log ID", http.StatusBadRequest)
		return
	}

	err = mealRepo.Delete(userID, mealID)
	if err != nil {
		http.Error(response, "Failed to delete meal log", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func UpdateMeal(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var meal models.Meal
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	mealID, err := strconv.ParseInt(vars["MealID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid meal  log ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&meal)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != meal.UserID || mealID != meal.MealID {
		http.Error(response, "Request payload does not match", http.StatusBadRequest)
		return
	}

	_, err = userRepo.GetUserByID(userID)
	if err != nil {
		if err.Error() != "User not found" {
			http.Error(response, "Failed to check user", http.StatusInternalServerError)
			return
		} else {
			http.Error(response, "User does not exist", http.StatusBadRequest)
			return
		}
	}

	_, err = mealRepo.Update(userID, mealID, &meal)
	if err != nil {
		http.Error(response, "Failed to update meal log", http.StatusInternalServerError)
		return
	}
	
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(meal)
}