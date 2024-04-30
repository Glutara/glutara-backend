package repository

import (
	"context"
	"time"

	"glutara/config"
	"glutara/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type BloodGlucoseLevelRepository interface {
	FindUserBloodGlucoseGraphicDataByDate(int64, time.Time) ([]models.GraphicData, error)
	Save(*models.BloodGlucoseLevel) (*models.BloodGlucoseLevel, error)
	GetUserBloodGlucoseLevelsMaxCount(int64) (int64, error)
}

type bloodGlucoseLevelRepo struct{}

func NewBloodGlucoseLevelRepository() BloodGlucoseLevelRepository {
	return &bloodGlucoseLevelRepo{}
}

func (*bloodGlucoseLevelRepo) FindUserBloodGlucoseGraphicDataByDate(userID int64, date time.Time) ([]models.GraphicData, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var glucoseGraphicDatas []models.GraphicData
	var glucoseGraphicData models.GraphicData
	nextDay := date.AddDate(0, 0, 1)

	itr := client.Collection(config.BloodGlucoseLevelCollection).Where("UserID", "==", userID).Where("Time", ">=", date).Where("Time", "<", nextDay).Select("Prediction", "Time").Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&glucoseGraphicData); err != nil {
			return nil, err
		}

		glucoseGraphicDatas = append(glucoseGraphicDatas, glucoseGraphicData)
	}

	return glucoseGraphicDatas, nil
}

func (*bloodGlucoseLevelRepo) Save(bloodGlucoseLevel *models.BloodGlucoseLevel) (*models.BloodGlucoseLevel, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.BloodGlucoseLevelCollection).Add(ctx, map[string]interface{}{
		"UserID":              bloodGlucoseLevel.UserID,
		"BloodGlucoseLevelID": bloodGlucoseLevel.BloodGlucoseLevelID,
		"Input":               bloodGlucoseLevel.Input,
		"Prediction":          bloodGlucoseLevel.Prediction,
		"Time":                bloodGlucoseLevel.Time,
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
