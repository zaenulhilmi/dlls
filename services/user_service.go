package services

import (
	"dlls/contracts"
)

func NewUserService(userRepo contracts.UserRepository) contracts.UserService {
	return &userServiceImpl{
		userRepository: userRepo,
	}
}

type userServiceImpl struct {
	userRepository contracts.UserRepository
}

// FindByID implements contracts.UserService.
func (u *userServiceImpl) FindByID(id string) (*contracts.User, error) {
	return u.userRepository.FindByID(id)
}

// Update implements contracts.UserService.
func (u *userServiceImpl) Update(id string, user contracts.User) error {
	return u.userRepository.Update(id, user)
}

