package services

import (
	"dlls/contracts"
)

func NewSubscriptionService(userRepository contracts.UserRepository) contracts.SubscriptionService {
	return &subscriptionServiceImpl{
		userRepository: userRepository,
	}
}

type subscriptionServiceImpl struct {
	userRepository contracts.UserRepository
}

func (s *subscriptionServiceImpl) Subscribe(userID string) error {
	if userID == "" {
		return contracts.ErrUserNotFound
	}

	user, err := s.userRepository.FindByID(userID)

	if err != nil {
		return err
	}

	if user == nil {
		return contracts.ErrUserNotFound
	}

	user.IsPremium = true

	return s.userRepository.Update(userID, *user)
}
