package contracts

type UserRepository interface {
	Save(user User) error
	FindByEmail(email string) (*User, error)
}
		

