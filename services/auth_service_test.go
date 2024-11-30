package services_test

import (
	"dlls/contracts"
	"dlls/repositories"
	"dlls/services"
	"dlls/utils"
	"testing"
)

func TestAuthService_Login_WrongEmail(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	authService := newAuthService(userRepo)
	_, err := authService.Login("email", "password")
	if err != contracts.ErrUserNotFound {
		t.Errorf("AuthService.Login() error = %v, wantErr %v", err, true)
	}
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	hasher := utils.NewHasher()

	user := contracts.User{
		Email:        "email",
		PasswordHash: hasher.Hash("password"),
	}
	userRepo := repositories.NewMemUserRepository()
	userRepo.Save(user)

	authService := newAuthService(userRepo)
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

	memUserRepository := repositories.NewMemUserRepository()
	memUserRepository.Save(user)

	authService := newAuthService(memUserRepository)

	token, err := authService.Login("email", "password")
	if err != nil {
		t.Errorf("AuthService.Login() error = %v, wantErr %v", err, false)
	}
	if token != expectedToken {
		t.Errorf("AuthService.Login() token = %v, want %v", token, expectedToken)
	}
}

func TestAuthService_SignUp_NoEmailAndPassword(t *testing.T) {
	userRepository := repositories.NewMemUserRepository()
	authService := newAuthService(userRepository)

	err := authService.SignUp("", "")
	if err == nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, true)
	}
}

func TestAuthService_SignUp_EmailExists(t *testing.T) {

	userRepo := repositories.NewMemUserRepository()
	userRepo.Save(contracts.User{
		Email: "email",
		PasswordHash: "password",
	})

	authService := newAuthService(userRepo)

	err := authService.SignUp("email", "password")

	if err == nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, true)
	}

	if err != contracts.ErrUserExists {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, "user already exists")
	}

}

func TestAuthService_SignUp_ValidUser(t *testing.T) {
	userRepo := repositories.NewMemUserRepository()
	authService := newAuthService(userRepo)

	err := authService.SignUp("email", "password")

	if err != nil {
		t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, false)
	}

}

func newAuthService(userRepository contracts.UserRepository) contracts.AuthService {
	hasher := utils.NewHasher()
	jwt := utils.NewJWT()
	return services.NewAuthService(userRepository, hasher, jwt)
}
