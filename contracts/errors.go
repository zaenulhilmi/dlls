package contracts

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrLoginFailed = errors.New("can not login")

var ErrInvalidToken = errors.New("invalid token")
var ErrTokenNotFound = errors.New("token not found")

var ErrEmptyName = errors.New("name can not be empty")
