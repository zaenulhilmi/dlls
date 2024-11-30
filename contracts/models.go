package contracts

type User struct {
	ID           string
	Name         string
	Email        string
	Phone        string
	PasswordHash string
}

type Action struct {
	ID         string
	UserID     string
	TargetID   string
	ActionType ActionType
	ActionedAt string
}

type ActionType string

const (
	ActionTypeLike ActionType = "like"
	ActionTypePass ActionType = "pass"
)
