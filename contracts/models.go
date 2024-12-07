package contracts

import "time"

type User struct {
	ID           string
	Name         string
	Email        string
	Phone        string
	PasswordHash string
	IsPremium    bool
}

type Action struct {
	ID         string
	UserID     string
	TargetID   string
	ActionType ActionType
	ActionedAt time.Time
}

type ActionType string

const (
	ActionTypeLike ActionType = "like"
	ActionTypePass ActionType = "pass"
)
