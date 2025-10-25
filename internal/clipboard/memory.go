package clipboard

import (
	"context"
	"sync"
	"time"
)

// InMemoryStore stores a single clipboard entry in memory using a RWMutex for synchronisation.
type InMemoryStore struct {
	mu    sync.RWMutex
	entry Entry
	empty bool
}

// NewInMemoryStore initialises a clipboard store without any content.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{empty: true}
}

// Set implements the Store interface.
func (s *InMemoryStore) Set(ctx context.Context, owner, content string) (Entry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.entry = Entry{
		Owner:     owner,
		Content:   content,
		UpdatedAt: time.Now().UTC(),
	}
	s.empty = false

	return s.entry, nil
}

// Get implements the Store interface.
func (s *InMemoryStore) Get(ctx context.Context) (Entry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.empty {
		return Entry{}, ErrEmpty
	}
	return s.entry, nil
}
