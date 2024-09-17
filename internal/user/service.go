package user

import (
	"nutripet/pkg/models"
)

type Service interface {
	RegisterUser(user *models.User) error
	Login(email, password string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) RegisterUser(user *models.User) error {
	// Aquí podrías agregar la lógica para cifrar la contraseña, validaciones, etc.
	return s.repo.CreateUser(user)
}

func (s *service) Login(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	// Validar la contraseña aquí
	return user, nil
}

func (s *service) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
