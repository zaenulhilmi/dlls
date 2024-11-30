package services_test

import (
	"dlls/contracts"
	"dlls/services"
	"dlls/utils"
	"testing"
)

func TestAuthService_Login_WrongEmail(t *testing.T) {
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

func TestAuthService_Login_WrongPassword(t *testing.T) {
	hasher := utils.NewHasher()

	userRepo := mockUserRepository{
		findByEmailResult: &contracts.User{
			Email:        "email",
			PasswordHash: hasher.Hash("password"),
		},
	}
	authService := newAuthService(&userRepo)
	token, err := authService.Login("email", "wrong-password")

	if err != contracts.ErrInvalidCredentials {
		t.Errorf("AuthService.Login() error = %v, wantErr %v", err, contracts.ErrInvalidCredentials)
	}

	if token != "" {
		t.Error("AuthService.Login() token is not empty")
	}

}

func TestAuthService_Login_ValidUser(t *testing.T) {
	hasher := utils.NewHasher()
	jwt := utils.NewJWT()

	user := contracts.User{
		Email:        "email",
		PasswordHash: hasher.Hash("password"),
	}

	expectedToken, _ := jwt.GenerateJWT(user)

	userRepo := mockUserRepository{
		findByEmailResult: &user,
	}
	authService := newAuthService(&userRepo)
	token, err := authService.Login("email", "password")
	if err != nil {
		t.Errorf("AuthService.Login() error = %v, wantErr %v", err, false)
	}
	if token != expectedToken {
		t.Errorf("AuthService.Login() token = %v, want %v", token, expectedToken)
	}
}

func TestAuthService_SignUp_NoEmailAndPassword(t *testing.T) {
	authService := newAuthService(&mockUserRepository{})

	err := authService.SignUp("", "", "")
	if err == nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, true)
	}
}

func TestAuthService_SignUp_EmailExists(t *testing.T) {
	userRepo := mockUserRepository{
		findByEmailResult: &contracts.User{
			Email: "email",
		},
	}

	authService := newAuthService(&userRepo)

	err := authService.SignUp("name", "email", "password")


	if err == nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, true)
	}

	if userRepo.isSaveCalled {
		t.Error("UserRepository.Save() was called")
	}

	if !userRepo.isFindByEmailCalled {
		t.Error("UserRepository.FindByEmail() was not called")
	}

	if err != contracts.ErrUserExists {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, "user already exists")
	}

}

func TestAuthService_SignUp_ValidUser(t *testing.T) {
	userRepo := mockUserRepository{}
	authService := newAuthService(&userRepo)

	err := authService.SignUp("name", "email", "password")

	if err != nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, false)
	}

	if !userRepo.isSaveCalled {
		t.Error("UserRepository.Save() was not called")
	}
}

func newAuthService(userRepository contracts.UserRepository) contracts.AuthService {
	hasher := utils.NewHasher()
	jwt := utils.NewJWT()
	return services.NewAuthService(userRepository, hasher, jwt)
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
