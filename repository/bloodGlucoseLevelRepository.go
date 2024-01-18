package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"glutara/models"
	"glutara/config"
)

type BloodGlucoseLevelRepository interface {
	FindAllUserBloodGlucoseLevels(int64) ([]models.BloodGlucoseLevel, error)
	Save(*models.BloodGlucoseLevel) (*models.BloodGlucoseLevel, error)
	GetUserBloodGlucoseLevelsMaxCount(int64) (int64, error)
}

type bloodGlucoseLevelRepo struct{}

func NewBloodGlucoseLevelRepository() BloodGlucoseLevelRepository {
	return &bloodGlucoseLevelRepo{}
}

func (*bloodGlucoseLevelRepo) FindAllUserBloodGlucoseLevels(userID int64) ([]models.BloodGlucoseLevel, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var bloodGlucoseLevels []models.BloodGlucoseLevel
	var bloodGlucoseLevel models.BloodGlucoseLevel

	itr := client.Collection(config.BloodGlucoseLevelCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&bloodGlucoseLevel); err != nil {
			return nil, err
		}

		bloodGlucoseLevels = append(bloodGlucoseLevels, bloodGlucoseLevel)
	}

	return bloodGlucoseLevels, nil
}

func (*bloodGlucoseLevelRepo) Save(bloodGlucoseLevel *models.BloodGlucoseLevel) (*models.BloodGlucoseLevel, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.BloodGlucoseLevelCollection).Add(ctx, map[string]interface{}{
		"UserID":				bloodGlucoseLevel.UserID,
		"BloodGlucoseLevelID":	bloodGlucoseLevel.BloodGlucoseLevelID,
		"Input":				bloodGlucoseLevel.Input,
		"Prediction":			bloodGlucoseLevel.Prediction,
		"Time": 				bloodGlucoseLevel.Time,
	})

	if err != nil {
		return nil, err
	}

	return bloodGlucoseLevel, nil
}

func (*bloodGlucoseLevelRepo) GetUserBloodGlucoseLevelsMaxCount(userID int64) (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return 0, err
	}

	defer client.Close()
	var maxID int64 = 0

	itr := client.Collection(config.BloodGlucoseLevelCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		
		if maxID < doc.Data()["BloodGlucoseLevelID"].(int64) {
			maxID = doc.Data()["BloodGlucoseLevelID"].(int64)
		}
	}

	return maxID, nil
}