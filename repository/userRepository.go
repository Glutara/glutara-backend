package repository

import (
	"context"
	"errors"

	"glutara/config"
	"glutara/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type UserRepository interface {
	Save(*models.User) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserByID(int64) (*models.User, error)
	GetUserCount() (int64, error)
	UpdateLatestBloodGlucose(int64, float32) (error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (*userRepo) Save(user *models.User) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.UserCollection).Add(ctx, map[string]interface{}{
		"ID":		user.ID,
		"Email":	user.Email,
		"Password":	user.Password,
		"Name":		user.Name,
		"Role":		user.Role,
		"Phone":	user.Phone,
		"LatestBloodGlucose":	user.LatestBloodGlucose,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (*userRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var user models.User

	itr := client.Collection(config.UserCollection).Where("Email", "==", email).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (*userRepo) GetUserByID(userID int64) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var user models.User

	itr := client.Collection(config.UserCollection).Where("ID", "==", userID).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (*userRepo) GetUserCount() (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return 0, err
	}

	defer client.Close()
	var count int64 = 0

	itr := client.Collection(config.UserCollection).Documents(ctx)
	for {
		_, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		count++
	}

	return count, nil
}

func (*userRepo) UpdateLatestBloodGlucose(userID int64, prediction float32) (error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return err
	}

	defer client.Close()

	itr := client.Collection(config.UserCollection).Where("ID", "==", userID).Documents(ctx)
	doc, err := itr.Next()
	if err != nil {
		return err
	}

	_, err = doc.Ref.Update(ctx, []firestore.Update{{Path: "LatestBloodGlucose", Value: prediction}})
	if err != nil {
		return err
	}

	return nil
}