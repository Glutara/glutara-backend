package repository

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"glutara/models"
	"glutara/config"
)

type MedicationRepository interface {
	FindAllUserMedications(int64) ([]models.Medication, error)
	Save(*models.Medication) (*models.Medication, error)
	GetUserMedicationsMaxCount(int64) (int64, error)
	Delete(int64, int64) error
	Update(int64, int64, *models.Medication) (*models.Medication, error)
}

type medicationRepo struct{}

func NewMedicationRepository() MedicationRepository {
	return &medicationRepo{}
}

func (*medicationRepo) FindAllUserMedications(userID int64) ([]models.Medication, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	var medications []models.Medication
	var medication models.Medication

	itr := client.Collection(config.MedicationCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err := doc.DataTo(&medication); err != nil {
			return nil, err
		}

		medications = append(medications, medication)
	}

	return medications, nil
}

func (*medicationRepo) Save(medication *models.Medication) (*models.Medication, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(config.MedicationCollection).Add(ctx, map[string]interface{}{
		"UserID":		medication.UserID,
		"MedicationID":	medication.MedicationID,
		"Type":			medication.Type,
		"Category":		medication.Category,
		"Dose":			medication.Dose,
		"Date":			medication.Date,
		"Time":			medication.Time,
	})

	if err != nil {
		return nil, err
	}

	return medication, nil
}

func (*medicationRepo) GetUserMedicationsMaxCount(userID int64) (int64, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return 0, err
	}

	defer client.Close()
	var maxID int64 = 0

	itr := client.Collection(config.MedicationCollection).Where("UserID", "==", userID).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		
		if maxID < doc.Data()["MedicationID"].(int64) {
			maxID = doc.Data()["MedicationID"].(int64)
		}
	}

	return maxID, nil
}

func (*medicationRepo) Delete(userID int64, medicationID int64) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return err
	}

	defer client.Close()

	itr := client.Collection(config.MedicationCollection).Where("UserID", "==", userID).Where("MedicationID", "==", medicationID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return err
	}

	if len(doc) != 1 {
		return errors.New("Medication log not found")
	}

	_, err = client.Collection(config.MedicationCollection).Doc(doc[0].Ref.ID).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (*medicationRepo) Update(userID int64, medicationID int64, newData *models.Medication) (*models.Medication, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, config.ProjectID)

	if err != nil {
		return nil, err
	}

	defer client.Close()
	itr := client.Collection(config.MedicationCollection).Where("UserID", "==", userID).Where("MedicationID", "==", medicationID).Documents(ctx)
	doc, err := itr.GetAll()
	if err != nil {
		return nil, err
	}

	if len(doc) != 1 {
		return nil, errors.New("Medication log not found")
	}

	_, err = doc[0].Ref.Set(ctx, *newData)
	if err != nil {
		return nil, err
	}

	return newData, nil
}