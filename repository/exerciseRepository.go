package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"glutara/models"
	"glutara/config"
)

type ExerciseRepository interface {
	FindAllUserExercises(int64) ([]models.Exercise, error)
	Save(*models.Exercise) (*models.Exercise, error)
	GetUserExercisesMaxCount(int64) (int64, error)
	Delete(int64, int64) error
	Update(int64, int64, *models.Exercise) (*models.Exercise, error)
}

type exerciseRepo struct{}

func NewExerciseRepository() ExerciseRepository {
	return &exerciseRepo{}
}

func (*exerciseRepo) FindAllUserExercises(userID int64) ([]models.Exercise, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var exercises []models.Exercise
	var exercise models.Exercise

	itr := client.Collection(config.ExerciseCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&exercise); err != nil {
			return nil, err
		}

		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

func (*exerciseRepo) Save(exercise *models.Exercise) (*models.Exercise, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.ExerciseCollection).Add(ctx, map[string]interface{}{
		"UserID":		exercise.UserID,
		"ExerciseID":	exercise.ExerciseID,
		"Name":			exercise.Name,
		"Intensity":	exercise.Intensity,
		"Date":			exercise.Date,
		"StartTime":	exercise.StartTime,
		"EndTime": 		exercise.EndTime,
	})

	if err != nil {
		return nil, err
	}

	return exercise, nil
}

func (*exerciseRepo) GetUserExercisesMaxCount(userID int64) (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return 0, err
	}

	defer client.Close()
	var maxID int64 = 0

	itr := client.Collection(config.ExerciseCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		
		if maxID < doc.Data()["ExerciseID"].(int64) {
			maxID = doc.Data()["ExerciseID"].(int64)
		}
	}

	return maxID, nil
}

func (*exerciseRepo) Delete(userID int64, exerciseID int64) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return err
	}

	defer client.Close()

	itr := client.Collection(config.ExerciseCollection).Where("UserID", "==", userID).Where("ExerciseID", "==", exerciseID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return err
	}

	if len(doc) != 1 {
		return errors.New("exercise log not found")
	}

	_, err = client.Collection(config.ExerciseCollection).Doc(doc[0].Ref.ID).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (*exerciseRepo) Update(userID int64, exerciseID int64, newData *models.Exercise) (*models.Exercise, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	itr := client.Collection(config.ExerciseCollection).Where("UserID", "==", userID).Where("ExerciseID", "==", exerciseID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return nil, err
	}

	if len(doc) != 1 {
		return nil, errors.New("exercise log not found")
	}

	_, err = doc[0].Ref.Set(ctx, *newData)
	if err != nil {
		return nil, err
	}

	return newData, nil
}