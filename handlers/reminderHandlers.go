package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"glutara/models"
	"glutara/repository"
)

func GetReminders(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	reminders, err := repository.ReminderRepo.FindAllUserReminders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reminders)
}

func CreateReminder(c *gin.Context) {
	var reminder models.Reminder

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&reminder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != reminder.UserID {
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

	max, err := repository.ReminderRepo.GetUserRemindersMaxCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}
	reminder.ReminderID = max + 1

	_, err = repository.ReminderRepo.Save(&reminder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusOK, reminder)
}

func DeleteReminder(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	reminderID, err := strconv.ParseInt(c.Param("ReminderID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = repository.ReminderRepo.Delete(userID, reminderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func UpdateReminder(c *gin.Context) {
	var reminder models.Reminder

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	reminderID, err := strconv.ParseInt(c.Param("ReminderID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&reminder)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != reminder.UserID || reminderID != reminder.ReminderID {
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

	_, err = repository.ReminderRepo.Update(userID, reminderID, &reminder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update data"})
		return
	}

	c.JSON(http.StatusOK, reminder)
}
