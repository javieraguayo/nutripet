package meal

import (
	"gorm.io/gorm"
	"nutripet/pkg/models"
)

type Repository interface {
	CreateMeal(meal *models.Meal) error
	GetMealsByPetID(petID uint) ([]models.Meal, error)
	GetNextMeal(petID uint) (*models.Meal, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateMeal(meal *models.Meal) error {
	return r.db.Create(meal).Error
}

func (r *repository) GetMealsByPetID(petID uint) ([]models.Meal, error) {
	var meals []models.Meal
	err := r.db.Where("pet_id = ?", petID).Find(&meals).Error
	return meals, err
}

func (r *repository) GetNextMeal(petID uint) (*models.Meal, error) {
	var meal models.Meal
	err := r.db.Where("pet_id = ?", petID).Order("time ASC").First(&meal).Error
	return &meal, err
}
