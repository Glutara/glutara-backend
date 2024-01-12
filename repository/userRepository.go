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

type UserRepository interface {
	Save(*models.User) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserCount() (int64, error)	
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (*userRepo) Save(user *models.User) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
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
	})

	if err != nil {
		log.Fatalf("Failed to add a new user: %v", err)
		return nil, err
	}

	return user, nil
}

func (*userRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var user models.User

	itr := client.Collection(config.UserCollection).Where("Email", "==", email).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		return nil, errors.New("User not found")
	}
	if err != nil {
		log.Fatalf("Failed to iterate the list of users: %v", err)
		return nil, err
	}

	if err := doc.DataTo(&user); err != nil {
		log.Fatalf("Failed to convert data to User struct: %v", err)
		return nil, err
	}

	return &user, nil
}

func (*userRepo) GetUserCount() (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
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
			log.Fatalf("Failed to iterate the list of users: %v", err)
			return 0, err
		}
		count++
	}

	return count, nil
}