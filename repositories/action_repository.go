package repositories

import (
	"dlls/contracts"
	"time"
)

func NewMemActionRepository() contracts.ActionRepository {
	return &memActionRepository{
		actions: []contracts.Action{},
	}
}

type memActionRepository struct {
	actions []contracts.Action
}

// GetUnactionedUserIDsToday implements contracts.ActionRepository.
func (m *memActionRepository) GetUnactionedUserIDsToday(userID string, limit int) ([]string, error) {

	var result []string

	for _, action := range m.actions {
		if action.UserID == userID {
			continue
		}

		if action.ActionedAt.Day() == time.Now().Day() {
			continue
		}

		result = append(result, action.UserID)
	}

	return result, nil
}

// FindByUserID implements contracts.ActionRepository.
func (m *memActionRepository) FindByUserID(userID string) ([]contracts.Action, error) {

	var actions []contracts.Action

	for _, action := range m.actions {
		if action.UserID == userID {
			actions = append(actions, action)
		}
	}

	return actions, nil
}

// Save implements contracts.ActionRepository.
func (m *memActionRepository) Save(action contracts.Action) error {
	m.actions = append(m.actions, action)
	return nil
}
