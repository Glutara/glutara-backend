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
	exerciseRepo repository.ExerciseRepository = repository.NewExerciseRepository()
)

func GetExercises(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	exercises, err := exerciseRepo.FindAllUserExercises(userID)
	if err != nil {
		http.Error(response, "Failed to retrieve exercise logs data", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(exercises)
}

func CreateExercise(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var exercise models.Exercise
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&exercise)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != exercise.UserID {
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

	max, err := exerciseRepo.GetUserExercisesMaxCount(userID)
	if err != nil {
		http.Error(response, "Failed to get max exercise log count", http.StatusInternalServerError)
		return
	}
	exercise.ExerciseID = max + 1

	_, err = exerciseRepo.Save(&exercise)
	if err != nil {
		http.Error(response, "Failed to create new exercise log", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(exercise)
}

func DeleteExercise(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	exerciseID, err := strconv.ParseInt(vars["ExerciseID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid exercise log ID", http.StatusBadRequest)
		return
	}

	err = exerciseRepo.Delete(userID, exerciseID)
	if err != nil {
		http.Error(response, "Failed to delete exercise log", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func UpdateExercise(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var exercise models.Exercise
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	exerciseID, err := strconv.ParseInt(vars["ExerciseID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid exercise  log ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&exercise)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != exercise.UserID || exerciseID != exercise.ExerciseID {
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

	_, err = exerciseRepo.Update(userID, exerciseID, &exercise)
	if err != nil {
		http.Error(response, "Failed to update exercise log", http.StatusInternalServerError)
		return
	}
	
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(exercise)
}