package contracts

import "net/http"

type Hasher interface {
	Hash(password string) string
	Compare(password, hashedPassword string) bool
}

type Jwt interface {
	GenerateJWT(user User) (string, error)
	ParseJWT(tokenString string) (*User, error)
	ExtractToken(r *http.Request) (string, error)
}
