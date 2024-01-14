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
	sleepRepo repository.SleepRepository = repository.NewSleepRepository()
)

func GetSleeps(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	sleeps, err := sleepRepo.FindAllUserSleeps(userID)
	if err != nil {
		http.Error(response, "Failed to retrieve sleep logs data", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(sleeps)
}

func CreateSleep(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var sleep models.Sleep
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&sleep)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != sleep.UserID {
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

	max, err := sleepRepo.GetUserSleepsMaxCount(userID)
	if err != nil {
		http.Error(response, "Failed to get max sleep log count", http.StatusInternalServerError)
		return
	}
	sleep.SleepID = max + 1

	_, err = sleepRepo.Save(&sleep)
	if err != nil {
		http.Error(response, "Failed to create new sleep log", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(sleep)
}

func DeleteSleep(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	sleepID, err := strconv.ParseInt(vars["SleepID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid sleep log ID", http.StatusBadRequest)
		return
	}

	err = sleepRepo.Delete(userID, sleepID)
	if err != nil {
		http.Error(response, "Failed to delete sleep log", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func UpdateSleep(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var sleep models.Sleep
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	sleepID, err := strconv.ParseInt(vars["SleepID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid sleep  log ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&sleep)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != sleep.UserID || sleepID != sleep.SleepID {
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

	_, err = sleepRepo.Update(userID, sleepID, &sleep)
	if err != nil {
		http.Error(response, "Failed to update sleep log", http.StatusInternalServerError)
		return
	}
	
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(sleep)
}