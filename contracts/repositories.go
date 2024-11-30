package contracts

type UserRepository interface {
	Save(user User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
	GetUsers(exludeIDs []string) ([]User, error)
}
		


type ActionRepository interface {
	Save(action Action) error
	FindByUserID(userID string) ([]Action, error)
}
