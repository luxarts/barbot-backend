package service

import (
	"barbot/internal/domain"
	"barbot/internal/repository"
)

type UsersService interface {
	Create(user domain.User) (*domain.UserDTO, error)
	GetByUUID(UUID string) *domain.UserDTO
}

type usersService struct {
	repo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) UsersService {
	return &usersService{repo: repo}
}

func (s *usersService) Create(user domain.User) (*domain.UserDTO, error) {
	userResponse, err := s.repo.Create(user)
	if userResponse != nil {
		return userResponse.ToDTO(), nil
	}
	return nil, err
}

func (s *usersService) GetByUUID(UUID string) *domain.UserDTO {
	user := s.repo.GetByUUID(UUID)

	if user == nil {
		return nil
	}

	return user.ToDTO()
}
