package middleware

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"bytes"
	"os"

	"glutara/repository"
	"glutara/models"
)

var (
	bloodGlucoseLevelRepo repository.BloodGlucoseLevelRepository = repository.NewBloodGlucoseLevelRepository()
)

func GetBloodGlucoseLevels(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	bloodGlucoseLevels, err := bloodGlucoseLevelRepo.FindAllUserBloodGlucoseLevels(userID)
	if err != nil {
		http.Error(response, "Failed to retrieve blood glucose levels data", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(bloodGlucoseLevels)
}

func CreateBloodGlucoseLevel(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var bloodGlucoseLevel models.BloodGlucoseLevel
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&bloodGlucoseLevel)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != bloodGlucoseLevel.UserID {
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

	max, err := bloodGlucoseLevelRepo.GetUserBloodGlucoseLevelsMaxCount(userID)
	if err != nil {
		http.Error(response, "Failed to get max blood glucose level count", http.StatusInternalServerError)
		return
	}
	bloodGlucoseLevel.BloodGlucoseLevelID = max + 1

	prediction, err := GetPrediction(bloodGlucoseLevel.Input)
	if err != nil {
		http.Error(response, "Failed to get blood glucose level prediction", http.StatusInternalServerError)
		return
	}
	bloodGlucoseLevel.Prediction = prediction

	_, err = bloodGlucoseLevelRepo.Save(&bloodGlucoseLevel)
	if err != nil {
		http.Error(response, "Failed to create new blood glucose level", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(bloodGlucoseLevel)
}

func envModelPortOr(port string) string {
	if envPort := os.Getenv("MODEL_SERVE_PORT"); envPort != "" {
		return ":" + envPort
	}

	return ":" + port
}

func GetPrediction(input float32) (float32, error) {
	modelServing := "http://localhost" + envModelPortOr("8605") + "/v1/models/glutara_model:predict"

	requestData := models.InferenceReqToken {
		Instances: []float32{input},
	}
	payload, err := json.Marshal(requestData)
	if err != nil {
		return 0, err
	}

	response, err := http.Post(modelServing, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	var responseData models.PredictionResToken
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return 0, err
	}

	return responseData.Predictions[0], nil
}