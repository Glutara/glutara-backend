package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"glutara/models"
	"glutara/config"
)

type PostRepository interface {
	Save(*models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}

type repo struct{}

// NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *models.Post) (*models.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(config.PostCollection).Add(ctx, map[string]interface{} {
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]models.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []models.Post

	itr := client.Collection(config.PostCollection).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := models.Post {
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	
	return posts, nil
}