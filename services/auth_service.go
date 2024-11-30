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
		return "", contracts.ErrUserNotFound
	}

	if !a.hasher.Compare(password, user.PasswordHash) {
		return "", contracts.ErrInvalidCredentials
	}

	return a.jwt.GenerateJWT(*user)

}

// SignUp implements contracts.AuthService.
func (a *authServiceImpl) SignUp(name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return errors.New("name, email and password are required")
	}

	existingUser, err := a.userRepository.FindByEmail(email)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return contracts.ErrUserExists
	}

	user := contracts.User{
		Name:         name,
		Email:        email,
		PasswordHash: a.hasher.Hash(password),
	}

	return a.userRepository.Save(user)
}
