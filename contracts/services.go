package contracts

type AuthService interface {
	SignUp(name, email, password string) error
	Login(email, password string) (string, error)
}
