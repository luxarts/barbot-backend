package service

import (
	"barbot/internal/defines"
	"barbot/internal/domain"
	"barbot/internal/repository"
	"barbot/internal/utils/jwt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type SessionsService interface {
	Create(dto domain.SessionDTO) *string
}
type sessionsService struct {
	usersRepo repository.UsersRepository
}

func NewSessionsService(usersRepo repository.UsersRepository) SessionsService {
	return &sessionsService{usersRepo: usersRepo}
}

func (s *sessionsService) Create(dto domain.SessionDTO) *string {
	user := s.usersRepo.GetByEmail(dto.Email)

	// Check if email exists
	if user == nil {
		return nil
	}

	// Check if password correct
	if user.Password != dto.Password {
		return nil
	}

	// Generate JWT
	secret := os.Getenv(defines.EnvJWTSecret)
	token := jwt.GenerateToken(jwt.Payload{Subject: user.UUID}, secret)

	return &token
}
