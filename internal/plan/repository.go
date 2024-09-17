package plan

import (
	"gorm.io/gorm"
	"nutripet/pkg/models"
)

type Repository interface {
	CreatePlan(plan *models.Plan) error
	GetPlansByPetID(petID uint) ([]models.Plan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreatePlan(plan *models.Plan) error {
	return r.db.Create(plan).Error
}

func (r *repository) GetPlansByPetID(petID uint) ([]models.Plan, error) {
	var plans []models.Plan
	err := r.db.Where("pet_id = ?", petID).Find(&plans).Error
	return plans, err
}
