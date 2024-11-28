package services

import (
	"dlls/contracts"
	"errors"
)

func NewAuthService(userRepository contracts.UserRepository) contracts.AuthService {
	return &authServiceImpl{
		userRepository: userRepository,
	}
}

type authServiceImpl struct {
	userRepository contracts.UserRepository
}

// Login implements contracts.AuthService.
func (a *authServiceImpl) Login(email string, password string) (string, error) {
	user, err := a.userRepository.FindByEmail(email) 

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}
	return "", errors.New("unimplemented")
}

// SignUp implements contracts.AuthService.
func (a *authServiceImpl) SignUp(user contracts.User) error {
	if user.Email == "" || user.PasswordHash == "" {
		return errors.New("email and password are required")
	}

	return a.userRepository.Save(user)
}
