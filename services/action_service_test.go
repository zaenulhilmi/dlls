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

	actionService := services.NewActionService(userRepo, actionRepo, 2)

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

	actionService := services.NewActionService(userRepo, actionRepo, 2)

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

	actionService := services.NewActionService(userRepo, actionRepo, 2)

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

	actionService := services.NewActionService(userRepo, actionRepo, 2)

	targetID, err := actionService.NextTarget("user-id")

	if err != nil {
		t.Errorf("ActionService.NextTarget() error = %v, wantErr %v", err, false)
	}

	if targetID != "" {
		t.Errorf("ActionService.NextTarget() targetId = %v, want %v", targetID, "")
	}
}

func TestActionService_Like_RegularUser(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "target-id-1",
		Email: "target-email-1",
	})
	userRepo.Save(contracts.User{
		ID:    "target-id-2",
		Email: "target-email-2",
	})

	userRepo.Save(contracts.User{
		ID:    "target-id-3",
		Email: "user-email-3",
	})

	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
	})

	// regular limit is 2
	actionService := services.NewActionService(userRepo, actionRepo, 2)

	targetID, _ := actionService.NextTarget("user-id")
	actionService.Like("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	actionService.Like("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	err := actionService.Like("user-id", targetID)

	if err != contracts.ErrActionLimitReached {
		t.Errorf("ActionService.Like() error = %v, wantErr %v", err, contracts.ErrActionLimitReached)
	}
}

func TestActionService_Pass_RegularUser(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "target-id-1",
		Email: "target-email-1",
	})
	userRepo.Save(contracts.User{
		ID:    "target-id-2",
		Email: "target-email-2",
	})

	userRepo.Save(contracts.User{
		ID:    "target-id-3",
		Email: "user-email-3",
	})

	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
	})

	// regular limit is 2
	actionService := services.NewActionService(userRepo, actionRepo, 2)

	targetID, _ := actionService.NextTarget("user-id")
	actionService.Pass("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	actionService.Pass("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	err := actionService.Pass("user-id", targetID)

	if err != contracts.ErrActionLimitReached {
		t.Errorf("ActionService.Pass() error = %v, wantErr %v", err, contracts.ErrActionLimitReached)
	}
}

func TestActionService_Like_PremiumUser(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "target-id-1",
		Email: "target-email-1",
	})
	userRepo.Save(contracts.User{
		ID:    "target-id-2",
		Email: "target-email-2",
	})

	userRepo.Save(contracts.User{
		ID:    "target-id-3",
		Email: "user-email-3",
	})

	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
		IsPremium: true,
	})

	// regular limit is 2
	actionService := services.NewActionService(userRepo, actionRepo, 2)

	targetID, _ := actionService.NextTarget("user-id")
	actionService.Like("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	actionService.Like("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	err := actionService.Like("user-id", targetID)

	if err != nil {
		t.Errorf("ActionService.Like() error = %v, wantErr %v", err, false)
	}
}

func TestActionService_Pass_PremiumUser(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	actionRepo := repositories.NewMemActionRepository()

	userRepo.Save(contracts.User{
		ID:    "target-id-1",
		Email: "target-email-1",
	})
	userRepo.Save(contracts.User{
		ID:    "target-id-2",
		Email: "target-email-2",
	})

	userRepo.Save(contracts.User{
		ID:    "target-id-3",
		Email: "user-email-3",
	})

	userRepo.Save(contracts.User{
		ID:    "user-id",
		Email: "user-email",
		IsPremium: true,
	})

	// regular limit is 2
	actionService := services.NewActionService(userRepo, actionRepo, 2)

	targetID, _ := actionService.NextTarget("user-id")
	actionService.Pass("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	actionService.Pass("user-id", targetID)
	targetID, _ = actionService.NextTarget("user-id")
	err := actionService.Pass("user-id", targetID)

	if err != nil {
		t.Errorf("ActionService.Pass() error = %v, wantErr %v", err, false)
	}
}
