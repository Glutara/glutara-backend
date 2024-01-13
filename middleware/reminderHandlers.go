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
	reminderRepo repository.ReminderRepository = repository.NewReminderRepository()
)

func GetReminders(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	reminders, err := reminderRepo.FindAllUserReminders(userID)
	if err != nil {
		http.Error(response, "Failed to retrieve reminders data", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(reminders)
}

func CreateReminder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var reminder models.Reminder
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&reminder)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != reminder.UserID {
		http.Error(response, "Request payload does not match", http.StatusBadRequest)
		return
	}

	max, err := reminderRepo.GetUserRemindersMaxCount(userID)
	if err != nil {
		http.Error(response, "Failed to get max reminder count", http.StatusInternalServerError)
		return
	}
	reminder.ReminderID = max + 1

	_, err = reminderRepo.Save(&reminder)
	if err != nil {
		http.Error(response, "Failed to create new reminder", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(reminder)
}

func DeleteReminder(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	reminderID, err := strconv.ParseInt(vars["ReminderID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	err = reminderRepo.Delete(userID, reminderID)
	if err != nil {
		http.Error(response, "Failed to delete reminder", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func UpdateReminder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var reminder models.Reminder
	vars := mux.Vars(request)
	userID, err := strconv.ParseInt(vars["UserID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid user ID", http.StatusBadRequest)
		return
	}
	reminderID, err := strconv.ParseInt(vars["ReminderID"], 10, 64)
	if err != nil {
		http.Error(response, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&reminder)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if userID != reminder.UserID || reminderID != reminder.ReminderID {
		http.Error(response, "Request payload does not match", http.StatusBadRequest)
		return
	}

	_, err = reminderRepo.Update(userID, reminderID, &reminder)
	if err != nil {
		http.Error(response, "Failed to update reminder", http.StatusInternalServerError)
		return
	}
	
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(reminder)
}