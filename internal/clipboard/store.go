package clipboard

import (
	"context"
	"errors"
	"time"
)

var (
	// ErrEmpty indicates that the clipboard does not contain any entry yet.
	ErrEmpty = errors.New("clipboard is empty")
)

// Entry represents a clipboard state at a given point in time.
type Entry struct {
	Owner     string    `json:"owner"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Store describes the behaviour that every clipboard persistence implementation must satisfy.
type Store interface {
	// Set updates the clipboard content for the specified owner within the provided scope.
	Set(ctx context.Context, scope, owner, content string) (Entry, error)
	// Get returns the latest clipboard entry for the provided scope.
	Get(ctx context.Context, scope string) (Entry, error)
}
