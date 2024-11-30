package services

import (
	"dlls/contracts"
	"math/rand/v2"
	"time"
)

func NewActionService(
	userRepository contracts.UserRepository,
	actionRepository contracts.ActionRepository,
	regularUserLimit int,
) contracts.ActionService {
	return &actionService{
		userRepository:   userRepository,
		actionRepository: actionRepository,
		regularUserLimit: regularUserLimit,
	}
}

type actionService struct {
	userRepository   contracts.UserRepository
	actionRepository contracts.ActionRepository
	regularUserLimit int
}

// Like implements contracts.ActionService.
func (a *actionService) Like(userID string, targetID string) error {
	return action(contracts.ActionTypeLike, userID, targetID, a)
}

// Pass implements contracts.ActionService.
func (a *actionService) Pass(userID string, targetID string) error {
	return action(contracts.ActionTypePass, userID, targetID, a)
}

func action(actionType contracts.ActionType, userID string, targetID string, a *actionService) error {
	action := contracts.Action{
		UserID:     userID,
		TargetID:   targetID,
		ActionType: actionType,
		ActionedAt: time.Now(),
	}

	user, err := a.userRepository.FindByID(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return contracts.ErrUserNotFound
	}


	if user.IsPremium {
		return a.actionRepository.Save(action)
	}

	actions, err := a.actionRepository.FindByUserID(userID)

	if err != nil {
		return err
	}

	actionsToday := getActionedUserIDsByUserIDToday(actions, userID)

	if(targetAlreadyActioned(actionsToday, targetID)) {
		return contracts.ErrActionAlreadyGiven
	}

	if len(actionsToday) >= a.regularUserLimit {
		return contracts.ErrActionLimitReached
	}

	return a.actionRepository.Save(action)
}

func targetAlreadyActioned(actionsToday []string, targetID string) bool {
	for _, actionedID := range actionsToday {
		if actionedID == targetID {
			return true
		}
	}
	return false
}

// NextTarget implements contracts.ActionService.
func (a *actionService) NextTarget(userID string) (string, error) {
	actions, err := a.actionRepository.FindByUserID(userID)
	if err != nil {
		return "", err
	}

	actionedUserIDs := getActionedUserIDsByUserIDToday(actions, userID)

	excludeIDs := append(actionedUserIDs, userID)
	users, err := a.userRepository.GetUsers(excludeIDs)

	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", nil
	}

	// random between 0 and len(users)

	randomIndex := randRange(0, len(users))


	return users[randomIndex].ID, nil
}

func getActionedUserIDsByUserIDToday(actions []contracts.Action, userID string) []string {
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

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
