package config

import (
	"os"
)

const ProjectID string = "upheld-acumen-420202"
const UserCollection string = "users"
const ReminderCollection string = "reminders"
const SleepCollection string = "sleeps"
const ExerciseCollection string = "exercises"
const MealCollection string = "meals"
const MedicationCollection string = "medications"
const BloodGlucoseLevelCollection string = "blood-glucose-levels"
const RelationCollection string = "relations"
var JWTKey []byte = []byte(envJWTSecretOr("glutara-gsc"))

func envJWTSecretOr(secret string) string {
	if envSecret := os.Getenv("JWT_SECRET"); envSecret != "" {
		return envSecret
	}
	return secret
}