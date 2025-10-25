package clipboard

import (
	"context"
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	store := NewInMemoryStore()
	ctx := context.Background()

	if _, err := store.Get(ctx); err != ErrEmpty {
		t.Fatalf("expected ErrEmpty, got %v", err)
	}

	entry, err := store.Set(ctx, "alice", "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entry.Owner != "alice" {
		t.Fatalf("owner mismatch: %s", entry.Owner)
	}

	got, err := store.Get(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Content != "hello" {
		t.Fatalf("content mismatch: %s", got.Content)
	}
}
