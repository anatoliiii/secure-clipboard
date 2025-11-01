package clipboard

import (
	"context"
	"strings"
	"sync"
	"time"
)

// DefaultScope is used when no scope is explicitly provided.
const DefaultScope = "shared"

// ScopedStore stores clipboard entries per scope in memory.
type ScopedStore struct {
	mu      sync.RWMutex
	entries map[string]Entry
}

// NewScopedStore initialises a clipboard store without any content.
func NewScopedStore() *ScopedStore {
	return &ScopedStore{entries: make(map[string]Entry)}
}

// Set implements the Store interface.
func (s *ScopedStore) Set(ctx context.Context, scope, owner, content string) (Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	normalizedScope := normalizeScope(scope)

	entry := Entry{
		Owner:     owner,
		Content:   content,
		UpdatedAt: time.Now().UTC(),
	}
	s.entries[normalizedScope] = entry

	return entry, nil
}

// Get implements the Store interface.
func (s *ScopedStore) Get(ctx context.Context, scope string) (Entry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	normalizedScope := normalizeScope(scope)

	entry, ok := s.entries[normalizedScope]
	if !ok {
		return Entry{}, ErrEmpty
	}
	return entry, nil
}

func normalizeScope(scope string) string {
	if strings.TrimSpace(scope) == "" {
		return DefaultScope
	}
	return scope
}
