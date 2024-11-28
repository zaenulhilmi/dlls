package repositories

import (
	"dlls/contracts"
)

func NewMemUserRepository() contracts.UserRepository {
	return &memUserRepository{
		users: make(map[string]contracts.User),
	}
}

type memUserRepository struct {
	users map[string]contracts.User
}

// FindByEmail implements contracts.UserRepository.
func (m *memUserRepository) FindByEmail(email string) (*contracts.User, error) {
	panic("unimplemented")
}

// Save implements contracts.UserRepository.
func (m *memUserRepository) Save(user contracts.User) error {
	panic("unimplemented")
}

