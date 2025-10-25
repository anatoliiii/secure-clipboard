package auth

import (
	"context"
	"crypto/subtle"
	"fmt"
	"strings"
)

// StaticAuthenticator implements Authenticator using a static in-memory map of users.
type StaticAuthenticator struct {
	users map[string]staticUser
}

type staticUser struct {
	password    string
	displayName string
}

// NewStaticAuthenticator creates an authenticator from a list of entries in the "username=password" or
// "username=password=Display Name" format.
func NewStaticAuthenticator(pairs []string) (*StaticAuthenticator, error) {
	users := make(map[string]staticUser)
	for _, pair := range pairs {
		if strings.TrimSpace(pair) == "" {
			continue
		}
		parts := strings.SplitN(pair, "=", 3)
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid credential pair %q: expected format username=password", pair)
		}
		username := strings.TrimSpace(parts[0])
		password := parts[1]
		displayName := username
		if len(parts) == 3 {
			displayName = parts[2]
		}
		users[username] = staticUser{password: password, displayName: displayName}
	}
	return &StaticAuthenticator{users: users}, nil
}

// Authenticate validates a username/password combination.
func (a *StaticAuthenticator) Authenticate(ctx context.Context, username, password string) (User, error) {
	if len(a.users) == 0 {
		return User{}, fmt.Errorf("no users configured: %w", ErrInvalidCredentials)
	}
	u, ok := a.users[username]
	if !ok {
		return User{}, ErrInvalidCredentials
	}
	if subtle.ConstantTimeCompare([]byte(u.password), []byte(password)) == 1 {
		return User{Username: username, DisplayName: u.displayName}, nil
	}
	return User{}, ErrInvalidCredentials
}
