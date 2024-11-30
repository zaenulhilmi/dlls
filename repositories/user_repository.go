package repositories

import (
	"dlls/contracts"
)

func NewMemUserRepository() contracts.UserRepository {
	return &memUserRepository{
		users: []contracts.User{},
	}
}

type memUserRepository struct {
	users []contracts.User
}

// Update implements contracts.UserRepository.
func (m *memUserRepository) Update(ID string, user contracts.User) error {
	for i, u := range m.users {
		if u.ID == ID {
			m.users[i] = user
			m.users[i].ID = ID
			return nil
		}
	}

	return contracts.ErrUserNotFound
}

// GetUsers implements contracts.UserRepository.
func (m *memUserRepository) GetUsers(exludeIDs []string) ([]contracts.User, error) {
	var result []contracts.User

	for _, user := range m.users {
		exclude := false
		for _, id := range exludeIDs {
			if user.ID == id {
				exclude = true
				break
			}
		}

		if !exclude {
			result = append(result, user)
		}
	}

	return result, nil
}

// FindByID implements contracts.UserRepository.
func (m *memUserRepository) FindByID(id string) (*contracts.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, nil
}

// FindByEmail implements contracts.UserRepository.
func (m *memUserRepository) FindByEmail(email string) (*contracts.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}

// Save implements contracts.UserRepository.
func (m *memUserRepository) Save(user contracts.User) error {
	for _, u := range m.users {
		if u.Email == user.Email {
			return contracts.ErrUserExists
		}
	}
	m.users = append(m.users, user)
	return nil
}
