package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"glutara/models"
	"glutara/config"
)

type SleepRepository interface {
	FindAllUserSleeps(int64) ([]models.Sleep, error)
	Save(*models.Sleep) (*models.Sleep, error)
	GetUserSleepsMaxCount(int64) (int64, error)
	Delete(int64, int64) error
	Update(int64, int64, *models.Sleep) (*models.Sleep, error)
}

type sleepRepo struct{}

func NewSleepRepository() SleepRepository {
	return &sleepRepo{}
}

func (*sleepRepo) FindAllUserSleeps(userID int64) ([]models.Sleep, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var sleeps []models.Sleep
	var sleep models.Sleep

	itr := client.Collection(config.SleepCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&sleep); err != nil {
			return nil, err
		}

		sleeps = append(sleeps, sleep)
	}

	return sleeps, nil
}

func (*sleepRepo) Save(sleep *models.Sleep) (*models.Sleep, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.SleepCollection).Add(ctx, map[string]interface{}{
		"UserID":		sleep.UserID,
		"SleepID":		sleep.SleepID,
		"StartTime":	sleep.StartTime,
		"EndTime": 		sleep.EndTime,
	})

	if err != nil {
		return nil, err
	}

	return sleep, nil
}

func (*sleepRepo) GetUserSleepsMaxCount(userID int64) (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return 0, err
	}

	defer client.Close()
	var maxID int64 = 0

	itr := client.Collection(config.SleepCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		
		if maxID < doc.Data()["SleepID"].(int64) {
			maxID = doc.Data()["SleepID"].(int64)
		}
	}

	return maxID, nil
}

func (*sleepRepo) Delete(userID int64, sleepID int64) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return err
	}

	defer client.Close()

	itr := client.Collection(config.SleepCollection).Where("UserID", "==", userID).Where("SleepID", "==", sleepID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return err
	}

	if len(doc) != 1 {
		return errors.New("Sleep log not found")
	}

	_, err = client.Collection(config.SleepCollection).Doc(doc[0].Ref.ID).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (*sleepRepo) Update(userID int64, sleepID int64, newData *models.Sleep) (*models.Sleep, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	itr := client.Collection(config.SleepCollection).Where("UserID", "==", userID).Where("SleepID", "==", sleepID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return nil, err
	}

	if len(doc) != 1 {
		return nil, errors.New("Sleep log not found")
	}

	_, err = doc[0].Ref.Set(ctx, *newData)
	if err != nil {
		return nil, err
	}

	return newData, nil
}