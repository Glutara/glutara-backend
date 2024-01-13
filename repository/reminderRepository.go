package repository

import (
	"context"
	"log"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"glutara/models"
	"glutara/config"
)

type ReminderRepository interface {
	FindAllUserReminders(int64) ([]models.Reminder, error)
	Save(*models.Reminder) (*models.Reminder, error)
	GetUserRemindersMaxCount(int64) (int64, error)
	Delete(int64, int64) error
	Update(int64, int64, *models.Reminder) (*models.Reminder, error)
}

type reminderRepo struct{}

func NewReminderRepository() ReminderRepository {
	return &reminderRepo{}
}

func (*reminderRepo) FindAllUserReminders(userID int64) ([]models.Reminder, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var reminders []models.Reminder
	var reminder models.Reminder

	itr := client.Collection(config.ReminderCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of reminders: %v", err)
			return nil, err
		}

		if err := doc.DataTo(&reminder); err != nil {
			log.Fatalf("Failed to convert data to Reminder struct: %v", err)
			return nil, err
		}

		reminders = append(reminders, reminder)
	}

	return reminders, nil
}

func (*reminderRepo) Save(reminder *models.Reminder) (*models.Reminder, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.ReminderCollection).Add(ctx, map[string]interface{}{
		"UserID":		reminder.UserID,
		"ReminderID":	reminder.ReminderID,
		"Name":			reminder.Name,
		"Description": 	reminder.Description,
		"Time":			reminder.Time,
	})

	if err != nil {
		log.Fatalf("Failed to add a new reminder: %v", err)
		return nil, err
	}

	return reminder, nil
}

func (*reminderRepo) GetUserRemindersMaxCount(userID int64) (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return 0, err
	}

	defer client.Close()
	var maxID int64 = 0

	itr := client.Collection(config.ReminderCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of reminders: %v", err)
			return 0, err
		}
		
		if maxID < doc.Data()["ReminderID"].(int64) {
			maxID = doc.Data()["ReminderID"].(int64)
		}
	}

	return maxID, nil
}

func (*reminderRepo) Delete(userID int64, reminderID int64) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return err
	}

	defer client.Close()

	itr := client.Collection(config.ReminderCollection).Where("UserID", "==", userID).Where("ReminderID", "==", reminderID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return err
	}

	if len(doc) != 1 {
		return errors.New("Reminder not found")
	}

	_, err = client.Collection(config.ReminderCollection).Doc(doc[0].Ref.ID).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (*reminderRepo) Update(userID int64, reminderID int64, newData *models.Reminder) (*models.Reminder, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	itr := client.Collection(config.ReminderCollection).Where("UserID", "==", userID).Where("ReminderID", "==", reminderID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return nil, err
	}

	if len(doc) != 1 {
		return nil, errors.New("Reminder not found")
	}

	_, err = doc[0].Ref.Set(ctx, *newData)
	if err != nil {
		return nil, err
	}

	return newData, nil
}