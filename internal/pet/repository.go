package pet

import (
	"gorm.io/gorm"
	"nutripet/pkg/models"
)

type Repository interface {
	CreatePet(pet *models.Pet) error
	GetPetsByUserID(userID uint) ([]models.Pet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreatePet(pet *models.Pet) error {
	return r.db.Create(pet).Error
}

func (r *repository) GetPetsByUserID(userID uint) ([]models.Pet, error) {
	var pets []models.Pet
	err := r.db.Where("user_id = ?", userID).Find(&pets).Error
	return pets, err
}
