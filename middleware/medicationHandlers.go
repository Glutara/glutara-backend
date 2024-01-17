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
	medicationRepo repository.MedicationRepository = repository.NewMedicationRepository()
)

func GetMedications(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	medications, err := medicationRepo.FindAllUserMedications(userID)
	if err != nil {
		http.Error(response, "Failed to retrieve medication logs data", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(medications)
}

func CreateMedication(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var medication models.Medication
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&medication)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != medication.UserID {
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

	max, err := medicationRepo.GetUserMedicationsMaxCount(userID)
	if err != nil {
		http.Error(response, "Failed to get max medication log count", http.StatusInternalServerError)
		return
	}
	medication.MedicationID = max + 1

	_, err = medicationRepo.Save(&medication)
	if err != nil {
		http.Error(response, "Failed to create new medication log", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(medication)
}

func DeleteMedication(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	medicationID, err := strconv.ParseInt(vars["MedicationID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid medication log ID", http.StatusBadRequest)
		return
	}

	err = medicationRepo.Delete(userID, medicationID)
	if err != nil {
		http.Error(response, "Failed to delete medication log", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func UpdateMedication(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var medication models.Medication
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	medicationID, err := strconv.ParseInt(vars["MedicationID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid medication  log ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&medication)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != medication.UserID || medicationID != medication.MedicationID {
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

	_, err = medicationRepo.Update(userID, medicationID, &medication)
	if err != nil {
		http.Error(response, "Failed to update medication log", http.StatusInternalServerError)
		return
	}
	
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(medication)
}