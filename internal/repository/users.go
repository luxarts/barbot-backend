package repository

import (
	"barbot/internal/defines"
	"barbot/internal/domain"

	"github.com/google/uuid"
)

type UsersRepository interface {
	Create(user domain.User) (*domain.User, error)
	GetByUUID(UUID string) *domain.User
	GetByEmail(email string) *domain.User
}

type usersRepository struct {
	db []domain.User
}

func NewUsersRepository() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) Create(user domain.User) (*domain.User, error) {
	if r.GetByEmail(user.Email) != nil {
		return nil, defines.ErrEmailAlreadyUsed
	}

	user.UUID = uuid.NewString()
	r.db = append(r.db, user)

	return &user, nil
}
func (r *usersRepository) GetByUUID(UUID string) *domain.User {
	for _, u := range r.db {
		if u.UUID == UUID {
			return &u
		}
	}

	return nil
}
func (r *usersRepository) GetByEmail(email string) *domain.User {
	for _, u := range r.db {
		if u.Email == email {
			return &u
		}
	}

	return nil
}
