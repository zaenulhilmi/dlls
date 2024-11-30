package services_test

import (
	"dlls/contracts"
	"dlls/repositories"
	"dlls/services"
	"testing"
)

func TestActionService_Like(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "target-id",
		Email: "target-email",
	})
	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
	})

	actionService := services.NewActionService(userRepo, actionRepo)

	err := actionService.Like("user-id", "target-id")

	if err != nil {
		t.Errorf("ActionService.Like() error = %v, wantErr %v", err, false)
	}

	actions, err := actionRepo.FindByUserID("user-id")

	if err != nil {
		t.Errorf("ActionService.Like() error = %v, wantErr %v", err, false)
	}

	if len(actions) != 1 {
		t.Errorf("ActionService.Like() len(actions) = %v, want %v", len(actions), 1)
	}

	if actions[0].UserID != "user-id" {
		t.Errorf("ActionService.Like() actions[0].UserID = %v, want %v", actions[0].UserID, "user-id")
	}

	if actions[0].ActionType != contracts.ActionTypeLike {
		t.Errorf("ActionService.Like() actions[0].ActionType = %v, want %v", actions[0].ActionType, contracts.ActionTypeLike)
	}
}

func TestActionService_Pass(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "target-id",
		Email: "target-email",
	})
	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
	})

	actionService := services.NewActionService(userRepo, actionRepo)

	err := actionService.Pass("user-id", "target-id")

	if err != nil {
		t.Errorf("ActionService.Pass() error = %v, wantErr %v", err, false)
	}

	actions, err := actionRepo.FindByUserID("user-id")

	if err != nil {
		t.Errorf("ActionService.Pass() error = %v, wantErr %v", err, false)
	}

	if len(actions) != 1 {
		t.Errorf("ActionService.Pass() len(actions) = %v, want %v", len(actions), 1)
	}

	if actions[0].UserID != "user-id" {
		t.Errorf("ActionService.Pass() actions[0].UserID = %v, want %v", actions[0].UserID, "user-id")
	}

	if actions[0].ActionType != contracts.ActionTypePass {
		t.Errorf("ActionService.Pass() actions[0].ActionType = %v, want %v", actions[0].ActionType, contracts.ActionTypePass)
	}
}

func TestActionService_NextTarget(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "target-id",
		Email: "target-email",
	})
	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
	})

	actionService := services.NewActionService(userRepo, actionRepo)

	targetID, err := actionService.NextTarget("user-id")

	if err != nil {
		t.Errorf("ActionService.NextTarget() error = %v, wantErr %v", err, false)
	}

	if targetID != "target-id" {
		t.Errorf("ActionService.NextTarget() targetId = %v, want %v", targetID, "target-id")
	}
}

func TestActionService_NextTarget_NoTarget(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
	})

	actionService := services.NewActionService(userRepo, actionRepo)

	targetID, err := actionService.NextTarget("user-id")

	if err != nil {
		t.Errorf("ActionService.NextTarget() error = %v, wantErr %v", err, false)
	}

	if targetID != "" {
		t.Errorf("ActionService.NextTarget() targetId = %v, want %v", targetID, "")
	}
}
