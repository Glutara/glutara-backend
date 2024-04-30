package repository

var (
	UserRepo UserRepository = NewUserRepository()
	BloodGlucoseLevelRepo BloodGlucoseLevelRepository = NewBloodGlucoseLevelRepository()
	ReminderRepo ReminderRepository = NewReminderRepository()
	MedicationRepo MedicationRepository = NewMedicationRepository()
	MealRepo MealRepository = NewMealRepository()
	SleepRepo SleepRepository = NewSleepRepository()
	ExerciseRepo ExerciseRepository = NewExerciseRepository()
	RelationRepo RelationRepository = NewRelationRepository()
)