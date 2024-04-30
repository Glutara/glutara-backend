package repository

import (
	"context"
	"errors"

	"glutara/config"
	"glutara/models"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type RelationRepository interface {
	Save(*models.Relation) (*models.Relation, error)
	GetAllUserRelations(int64) ([]models.Relation, error)
	GetAllUserRelatedInfos(int64) ([]models.RelatedInfo, error)
	CheckRelationExist(int64, int64) (error)
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

func (*relationRepo) GetAllUserRelatedInfos(relationID int64) ([]models.RelatedInfo, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var related models.Related
	var relatedInfos []models.RelatedInfo

	itr := client.Collection(config.RelationCollection).Where("RelationID", "==", relationID).Select("UserID", "Name", "Phone").Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&related); err != nil {
			return nil, err
		}

		newItr := client.Collection(config.UserCollection).Where("ID", "==", related.UserID).Select("LatestBloodGlucose").Documents(ctx)
		newDoc, err := newItr.Next()
		if err != nil {
			return nil, err
		}

		var latestBloodGlucoseInfo models.LatestBloodGlucoseInfo
		if err := newDoc.DataTo(&latestBloodGlucoseInfo); err != nil {
			return nil, err
		}

		relatedInfos = append(relatedInfos, models.RelatedInfo{
			UserID: related.UserID,
			Name: related.Name,
			Phone: related.Phone,
			LatestBloodGlucose: latestBloodGlucoseInfo.LatestBloodGlucose,
		})
	}

	return relatedInfos, nil
}

func (*relationRepo) CheckRelationExist(userID int64, relationID int64) (error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return err
	}

	defer client.Close()

	itr := client.Collection(config.RelationCollection).Where("UserID", "==", userID).Where("RelationID", "==", relationID).Documents(ctx)
	_, err = itr.Next()
	if err == iterator.Done {
		return errors.New("relation not found")
	}
	if err != nil {
		return err
	}
	return nil
}
