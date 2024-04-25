package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"glutara/models"
	"glutara/repository"
)

func GetMeals(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	meals, err := repository.MealRepo.FindAllUserMeals(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve data"})
		return
	}

	c.JSON(http.StatusOK, meals)
}

func CreateMeal(c *gin.Context) {
	var meal models.Meal

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&meal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != meal.UserID {
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

	max, err := repository.MealRepo.GetUserMealsMaxCount(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}
	meal.MealID = max + 1

	_, err = repository.MealRepo.Save(&meal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusOK, meal)
}

func DeleteMeal(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	mealID, err := strconv.ParseInt(c.Param("MealID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = repository.MealRepo.Delete(userID, mealID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func UpdateMeal(c *gin.Context) {
	var meal models.Meal

	userID, err := strconv.ParseInt(c.Param("UserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}
	mealID, err := strconv.ParseInt(c.Param("MealID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	err = c.BindJSON(&meal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request payload"})
		return
	}

	if userID != meal.UserID || mealID != meal.MealID {
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

	_, err = repository.MealRepo.Update(userID, mealID, &meal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update data"})
		return
	}

	c.JSON(http.StatusOK, meal)
}
