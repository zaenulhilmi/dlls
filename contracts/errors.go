package contracts

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrUserExists = errors.New("user already exists")

var ErrLoginFailed = errors.New("can not login")

var ErrInvalidToken = errors.New("invalid token")
var ErrTokenNotFound = errors.New("token not found")

var ErrEmptyName = errors.New("name can not be empty")

var ErrInvalidCredentials = errors.New("email or password is incorrect")

var ErrActionLimitReached = errors.New("action limit reached")
var ErrActionAlreadyGiven = errors.New("action already given, please preview next profile")
