package contracts

type AuthService interface {
	SignUp(name, email, password string) error
	Login(email, password string) (string, error)
}

type ActionService interface {
	Like(userID, targetID string) error
	Pass(userID, targetID string) error
	NextTarget(userID string) (string, error)
}

type SubscriptionService interface {
	Subscribe(userID string) error
}
