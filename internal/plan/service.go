package plan

import (
	"nutripet/pkg/models"
)

type Service interface {
	AddPlan(plan *models.Plan) error
	GetPetPlans(petID uint) ([]models.Plan, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) AddPlan(plan *models.Plan) error {
	return s.repo.CreatePlan(plan)
}

func (s *service) GetPetPlans(petID uint) ([]models.Plan, error) {
	return s.repo.GetPlansByPetID(petID)
}
