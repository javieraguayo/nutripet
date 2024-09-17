package pet

import (
	"nutripet/pkg/models"
)

type Service interface {
	AddPet(pet *models.Pet) error
	GetUserPets(userID uint) ([]models.Pet, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) AddPet(pet *models.Pet) error {
	return s.repo.CreatePet(pet)
}

func (s *service) GetUserPets(userID uint) ([]models.Pet, error) {
	return s.repo.GetPetsByUserID(userID)
}
