package contracts

type AuthService interface {
	SignUp(user User) error
	Login(email, password string) (string, error)
}
