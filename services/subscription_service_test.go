package services_test

import (
	"dlls/contracts"
	"dlls/services"
	"dlls/repositories"
	"testing"
)

func TestSubscriptionService_Subscribe(t *testing.T) {

	userRepo := repositories.NewMemUserRepository()
	subscriptionService := services.NewSubscriptionService(userRepo)

	user := contracts.User{
		ID:    "user-id",
		Email: "user-email",
	}

	userRepo.Save(user)

	err := subscriptionService.Subscribe("user-id")

	if err != nil {
		t.Errorf("SubscriptionService.Subscribe() error = %v, wantErr %v", err, false)
	}

	users, _ := userRepo.GetUsers([]string{})

	user = users[0]

	if !user.IsPremium {
		t.Error("SubscriptionService.Subscribe() user is not premium")
	}
}

func TestSubscriptionService_Subscribe_InvalidUserID(t *testing.T) {

	userRepo := repositories.NewMemUserRepository()
	subscriptionService := services.NewSubscriptionService(userRepo)

	err := subscriptionService.Subscribe("")

	if err != contracts.ErrUserNotFound {
		t.Errorf("SubscriptionService.Subscribe() error = %v, wantErr %v", err, contracts.ErrUserNotFound)
	}

}
