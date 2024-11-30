package services

import (
	"dlls/contracts"
)

func NewActionService() contracts.ActionService {
	return &actionService{}
}

type actionService struct{}

// Actions implements contracts.ActionService.
func (a *actionService) Actions(userID string) ([]contracts.Action, error) {
	panic("unimplemented")
}

// Like implements contracts.ActionService.
func (a *actionService) Like(userID string, targetID string) error {
	return nil
}

// NextTarget implements contracts.ActionService.
func (a *actionService) NextTarget(userID string) (string, error) {
	panic("unimplemented")
}

// Pass implements contracts.ActionService.
func (a *actionService) Pass(userID string, targetID string) error {
	panic("unimplemented")
}
