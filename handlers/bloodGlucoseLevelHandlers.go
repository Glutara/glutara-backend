package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"glutara/models"
	"glutara/repository"
)

func GetBloodGlucoseLevels(c * gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	bloodGlucoseLevels, err := repository.BloodGlucoseLevelRepo.FindAllUserBloodGlucoseLevels(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve data"})
		return
	}

	c.JSON(http.StatusOK, bloodGlucoseLevels)
}

func CreateBloodGlucoseLevel(c * gin.Context) {
	var bloodGlucoseLevel models.BloodGlucoseLevel

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&bloodGlucoseLevel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != bloodGlucoseLevel.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	_, err = repository.UserRepo.GetUserByID(userID)
	if err != nil {
		if err.Error() != "user not found" {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
			return
		}
	}

	max, err := repository.BloodGlucoseLevelRepo.GetUserBloodGlucoseLevelsMaxCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}
	bloodGlucoseLevel.BloodGlucoseLevelID = max + 1

	prediction, err := GetPrediction(bloodGlucoseLevel.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to predict data"})
		return
	}
	bloodGlucoseLevel.Prediction = prediction

	_, err = repository.BloodGlucoseLevelRepo.Save(&bloodGlucoseLevel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusOK, bloodGlucoseLevel)
}

func envModelPortOr(port string) string {
	if envPort := os.Getenv("MODEL_SERVE_PORT"); envPort != "" {
		return ":" + envPort
	}

	return ":" + port
}

func GetPrediction(input float32) (float32, error) {
	modelServing := "http://localhost" + envModelPortOr("8605") + "/v1/models/glutara_model:predict"

	requestData := models.InferenceReqToken{
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
