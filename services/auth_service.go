package services

import (
	"dlls/contracts"
	"errors"
)

func NewAuthService(
	userRepository contracts.UserRepository,
	hasher contracts.Hasher,
	jwt contracts.JWT,
) contracts.AuthService {
	return &authServiceImpl{
		userRepository: userRepository,
		hasher:         hasher,
		jwt:            jwt,
	}
}

type authServiceImpl struct {
	userRepository contracts.UserRepository
	hasher         contracts.Hasher
	jwt            contracts.JWT
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

	if !a.hasher.Compare(password, user.PasswordHash) {
		return "", contracts.ErrInvalidCredentials
	}

	return a.jwt.GenerateJWT(*user)

}

// SignUp implements contracts.AuthService.
func (a *authServiceImpl) SignUp(user contracts.User) error {
	if user.Email == "" || user.PasswordHash == "" {
		return errors.New("email and password are required")
	}

	return a.userRepository.Save(user)
}
