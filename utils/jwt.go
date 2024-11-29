package utils

import (
	"dlls/contracts"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJWT() contracts.JWT {
	return &jwtImpl{}
}

type jwtImpl struct{}

// ExtractToken implements contracts.Jwt.
func (j *jwtImpl) ExtractToken(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		return "", contracts.ErrTokenNotFound
	}

	tokenString = tokenString[7:]

	return tokenString, nil
}

// GenerateJWT implements contracts.Jwt.
func (j *jwtImpl) GenerateJWT(user contracts.User) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claimds := jwt.MapClaims{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"phone":   user.Phone,
		"exp":     expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimds)

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT implements contracts.Jwt.
func (j *jwtImpl) ParseJWT(tokenString string) (*contracts.User, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, contracts.ErrInvalidToken
	}

	user := contracts.User{
		ID:    claims["user_id"].(string),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
		Phone: claims["phone"].(string),
	}

	return &user, nil
}
