package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"glutara/models"
	"glutara/config"
)

type MealRepository interface {
	FindAllUserMeals(int64) ([]models.Meal, error)
	Save(*models.Meal) (*models.Meal, error)
	GetUserMealsMaxCount(int64) (int64, error)
	Delete(int64, int64) error
	Update(int64, int64, *models.Meal) (*models.Meal, error)
}

type mealRepo struct{}

func NewMealRepository() MealRepository {
	return &mealRepo{}
}

func (*mealRepo) FindAllUserMeals(userID int64) ([]models.Meal, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var meals []models.Meal
	var meal models.Meal

	itr := client.Collection(config.MealCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&meal); err != nil {
			return nil, err
		}

		meals = append(meals, meal)
	}

	return meals, nil
}

func (*mealRepo) Save(meal *models.Meal) (*models.Meal, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.MealCollection).Add(ctx, map[string]interface{}{
		"UserID":		meal.UserID,
		"MealID":		meal.MealID,
		"Name":			meal.Name,
		"Calories":		meal.Calories,
		"Type":			meal.Type,
		"Date":			meal.Date,
		"StartTime":	meal.StartTime,
		"EndTime": 		meal.EndTime,
	})

	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (*mealRepo) GetUserMealsMaxCount(userID int64) (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return 0, err
	}

	defer client.Close()
	var maxID int64 = 0

	itr := client.Collection(config.MealCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		
		if maxID < doc.Data()["MealID"].(int64) {
			maxID = doc.Data()["MealID"].(int64)
		}
	}

	return maxID, nil
}

func (*mealRepo) Delete(userID int64, mealID int64) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return err
	}

	defer client.Close()

	itr := client.Collection(config.MealCollection).Where("UserID", "==", userID).Where("MealID", "==", mealID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return err
	}

	if len(doc) != 1 {
		return errors.New("Meal log not found")
	}

	_, err = client.Collection(config.MealCollection).Doc(doc[0].Ref.ID).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (*mealRepo) Update(userID int64, mealID int64, newData *models.Meal) (*models.Meal, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	itr := client.Collection(config.MealCollection).Where("UserID", "==", userID).Where("MealID", "==", mealID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return nil, err
	}

	if len(doc) != 1 {
		return nil, errors.New("Meal log not found")
	}

	_, err = doc[0].Ref.Set(ctx, *newData)
	if err != nil {
		return nil, err
	}

	return newData, nil
}