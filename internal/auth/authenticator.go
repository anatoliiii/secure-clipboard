package auth

import (
	"context"
	"errors"
)

var (
	// ErrInvalidCredentials is returned when the provided credentials do not match a known user.
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// User represents an authenticated entity.
type User struct {
	Username    string
	DisplayName string
}

// Authenticator is responsible for verifying provided credentials and returning a User.
type Authenticator interface {
	Authenticate(ctx context.Context, username, password string) (User, error)
}
