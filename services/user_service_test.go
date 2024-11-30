package services_test

import (
	"dlls/contracts"
	"dlls/repositories"
	"dlls/services"
	"testing"
)

func TestUserService_FindByID(t *testing.T) {

	userRepo := repositories.NewMemUserRepository()
	userRepo.Save(contracts.User{
		ID:    "id",
		Email: "email",
	})
	userService := services.NewUserService(userRepo)

	user, err := userService.FindByID("id")

	if err != nil {
		t.Errorf("UserService.FindByID() error = %v, wantErr %v", err, false)
	}

	if user == nil {
		t.Error("UserService.FindByID() user is nil")
	}
}

func TestUserService_Update(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	userRepo.Save(contracts.User{
		ID:    "id",
		Email: "email",
	})
	userService := services.NewUserService(userRepo)

	err := userService.Update("id", contracts.User{
		Email: "new-email",
	})

	if err != nil {
		t.Errorf("UserService.Update() error = %v, wantErr %v", err, false)
	}

	user, err := userRepo.FindByID("id")

	if user.Email != "new-email" {
		t.Errorf("UserService.Update() user.Email = %v, want %v", user.Email, "new-email")
	}
}
