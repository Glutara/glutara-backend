package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"glutara/models"
	"glutara/repository"
)

func GetLogs(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	sleeps, err := repository.SleepRepo.FindAllUserSleeps(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve data"})
		return
	}

	meals, err := repository.MealRepo.FindAllUserMeals(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve data"})
		return
	}

	exercises, err := repository.ExerciseRepo.FindAllUserExercises(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve data"})
		return
	}

	medications, err := repository.MedicationRepo.FindAllUserMedications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve data"})
		return
	}

	var logs []models.Log
	for _, sleep := range sleeps {
		logs = append(logs, models.SleepLog{Type: "Sleep", Data: sleep})
	}
	for _, meal := range meals {
		logs = append(logs, models.MealLog{Type: "Meal", Data: meal})
	}
	for _, exercise := range exercises {
		logs = append(logs, models.ExerciseLog{Type: "Exercise", Data: exercise})
	}
	for _, medication := range medications {
		logs = append(logs, models.MedicationLog{Type: "Medication", Data: medication})
	}

	sort.Slice(logs, func (i, j int) bool {
		return logs[i].GetTime().After(logs[j].GetTime())
	})
	
	i := -1
	var groupedLogs []models.DateLogs
	for _, log := range logs {
		date := log.GetTime().Truncate(24 * time.Hour)
		if isDateNotAdded(date, groupedLogs) {
			groupedLogs = append(groupedLogs, models.DateLogs{Date: date, Logs: make([]models.Log, 0)})
			i += 1
		}
		groupedLogs[i].Logs = append(groupedLogs[i].Logs, log)
	}

	c.JSON(http.StatusOK, groupedLogs)
}

func isDateNotAdded(date time.Time, groupedLogs []models.DateLogs) bool {
	notHere := true

	for _, datelogs := range groupedLogs {
		if date.Equal(datelogs.Date) {
			notHere = false
			break
		}
	}

	return notHere
}
