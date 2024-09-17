package meal

import (
	"nutripet/pkg/models"
)

type Service interface {
	AddMeal(meal *models.Meal) error
	GetPetMeals(petID uint) ([]models.Meal, error)
	GetNextMeal(petID uint) (*models.Meal, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) AddMeal(meal *models.Meal) error {
	return s.repo.CreateMeal(meal)
}

func (s *service) GetPetMeals(petID uint) ([]models.Meal, error) {
	return s.repo.GetMealsByPetID(petID)
}

func (s *service) GetNextMeal(petID uint) (*models.Meal, error) {
	return s.repo.GetNextMeal(petID)
}
