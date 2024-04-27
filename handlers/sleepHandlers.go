package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"glutara/models"
	"glutara/repository"
)

func GetSleeps(c *gin.Context) {
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

	c.JSON(http.StatusOK, sleeps)
}

func CreateSleep(c *gin.Context) {
	var sleep models.Sleep

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&sleep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != sleep.UserID {
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

	max, err := repository.SleepRepo.GetUserSleepsMaxCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}
	sleep.SleepID = max + 1

	_, err = repository.SleepRepo.Save(&sleep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusOK, sleep)
}

func DeleteSleep(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	sleepID, err := strconv.ParseInt(c.Param("SleepID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = repository.SleepRepo.Delete(userID, sleepID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func UpdateSleep(c *gin.Context) {
	var sleep models.Sleep

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	sleepID, err := strconv.ParseInt(c.Param("SleepID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&sleep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != sleep.UserID || sleepID != sleep.SleepID {
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

	_, err = repository.SleepRepo.Update(userID, sleepID, &sleep)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update data"})
		return
	}

	c.JSON(http.StatusOK, sleep)
}
