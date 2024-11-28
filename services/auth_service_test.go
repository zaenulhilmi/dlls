package services_test

import (
	"dlls/contracts"
	"dlls/services"
	"testing"
)

func TestAuthService_Login_WrongEmailOrPassword(t *testing.T) {
	userRepo := mockUserRepository{
		findByEmailResult: nil,
	}
	authService := newAuthService(&userRepo)
	_, err := authService.Login("email", "password")
	if err == nil {
		t.Errorf("AuthService.Login() error = %v, wantErr %v", err, true)
	}
	if !userRepo.isFindByEmailCalled {
		t.Error("UserRepository.FindByEmail() was not called")
	}
}

func TestAuthService_SignUp_NoEmailAndPassword(t *testing.T) {
	authService := newAuthService(&mockUserRepository{})

	err := authService.SignUp(contracts.User{})
	if err == nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, true)
	}
}

func TestAuthService_SignUp_ValidUser(t *testing.T) {
	userRepo := mockUserRepository{}
	authService := newAuthService(&userRepo)

	err := authService.SignUp(contracts.User{
		Email:        "email",
		PasswordHash: "password",
	})

	if err != nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, false)
	}

	if !userRepo.isSaveCalled {
		t.Error("UserRepository.Save() was not called")
	}
}

func newAuthService(userRepository contracts.UserRepository) contracts.AuthService {
	return services.NewAuthService(userRepository)
}

type mockUserRepository struct {
	isSaveCalled        bool
	isFindByEmailCalled bool
	findByEmailResult   *contracts.User
	saveError           error
}

func (m *mockUserRepository) FindByEmail(email string) (*contracts.User, error) {
	m.isFindByEmailCalled = true
	return m.findByEmailResult, nil
}

func (m *mockUserRepository) Save(user contracts.User) error {
	m.isSaveCalled = true
	return m.saveError
}
