package services

import (
	"dlls/contracts"
	"time"
)

func NewActionService(userRepository contracts.UserRepository, actionRepository contracts.ActionRepository) contracts.ActionService {
	return &actionService{
		userRepository:   userRepository,
		actionRepository: actionRepository,
	}
}

type actionService struct {
	userRepository   contracts.UserRepository
	actionRepository contracts.ActionRepository
}

// Like implements contracts.ActionService.
func (a *actionService) Like(userID string, targetID string) error {
	action := contracts.Action{
		UserID:     userID,
		TargetID:   targetID,
		ActionType: contracts.ActionTypeLike,
		ActionedAt: time.Now(),
	}

	return a.actionRepository.Save(action)
}

// NextTarget implements contracts.ActionService.
func (a *actionService) NextTarget(userID string) (string, error) {
	actions, err := a.actionRepository.FindByUserID(userID)
	if err != nil {
		return "", err
	}

	actionedUserIDs := getActionedUserIDsByUserID(actions, userID)

	excludeIDs := append(actionedUserIDs, userID)
	users, err := a.userRepository.GetUsers(excludeIDs)

	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", nil
	}

	return users[0].ID, nil
}

// Pass implements contracts.ActionService.
func (a *actionService) Pass(userID string, targetID string) error {
	action := contracts.Action{
		UserID:     userID,
		TargetID:   targetID,
		ActionType: contracts.ActionTypePass,
		ActionedAt: time.Now(),
	}

	return a.actionRepository.Save(action)
}

func getActionedUserIDsByUserID(actions []contracts.Action, userID string) []string {
	var result []string

	for _, action := range actions {
		if action.UserID != userID {
			continue
		}

		if action.ActionedAt.Day() != time.Now().Day() {
			continue
		}

		result = append(result, action.TargetID)
	}

	return result
}
