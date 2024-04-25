package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"glutara/models"
	"glutara/repository"
)

func GetMedications(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	medications, err := repository.MedicationRepo.FindAllUserMedications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve data"})
		return
	}

	c.JSON(http.StatusOK, medications)
}

func CreateMedication(c *gin.Context) {
	var medication models.Medication

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&medication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != medication.UserID {
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

	max, err := repository.MedicationRepo.GetUserMedicationsMaxCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}
	medication.MedicationID = max + 1

	_, err = repository.MedicationRepo.Save(&medication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusOK, medication)
}

func DeleteMedication(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	medicationID, err := strconv.ParseInt(c.Param("MedicationID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = repository.MedicationRepo.Delete(userID, medicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func UpdateMedication(c *gin.Context) {
	var medication models.Medication

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	medicationID, err := strconv.ParseInt(c.Param("MedicationID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&medication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != medication.UserID || medicationID != medication.MedicationID {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	_, err = repository.UserRepo.GetUserByID(userID)
	if err != nil {
		if err.Error() != "user not found" {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update data"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
			return
		}
	}

	_, err = repository.MedicationRepo.Update(userID, medicationID, &medication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update data"})
		return
	}

	c.JSON(http.StatusOK, medication)
}
