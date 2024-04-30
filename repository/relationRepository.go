package repository

import (
	"context"

	"glutara/config"
	"glutara/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type RelationRepository interface {
	Save(*models.Relation) (*models.Relation, error)
	GetAllUserRelations(int64) ([]models.Relation, error)
}

type relationRepo struct{}

func NewRelationRepository() RelationRepository {
	return &relationRepo{}
}

func (*relationRepo) Save(relation *models.Relation) (*models.Relation, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.RelationCollection).Add(ctx, map[string]interface{}{
		"UserID":     relation.UserID,
		"Name":       relation.Name,
		"Phone":      relation.Phone,
		"RelationID": relation.RelationID,
		"RelationName": relation.RelationName,
		"RelationPhone": relation.RelationPhone,
		"Longitude":  relation.Longitude,
		"Latitude":   relation.Latitude,
	})

	if err != nil {
		return nil, err
	}

	return relation, nil
}

func (*relationRepo) GetAllUserRelations(userID int64) ([]models.Relation, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var relations []models.Relation
	var relation models.Relation

	itr := client.Collection(config.RelationCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&relation); err != nil {
			return nil, err
		}

		relations = append(relations, relation)
	}

	return relations, nil
}
