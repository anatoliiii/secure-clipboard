package clipboard

import (
	"context"
	"testing"
)

func TestScopedStore_DefaultScope(t *testing.T) {
	store := NewScopedStore()
	ctx := context.Background()

	if _, err := store.Get(ctx, ""); err != ErrEmpty {
		t.Fatalf("expected ErrEmpty, got %v", err)
	}

	if _, err := store.Get(ctx, "shared"); err != ErrEmpty {
		t.Fatalf("expected ErrEmpty for shared scope, got %v", err)
	}

	entry, err := store.Set(ctx, "", "alice", "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entry.Owner != "alice" {
		t.Fatalf("owner mismatch: %s", entry.Owner)
	}

	got, err := store.Get(ctx, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Content != "hello" {
		t.Fatalf("content mismatch: %s", got.Content)
	}
}

func TestScopedStore_IsolatedScopes(t *testing.T) {
	store := NewScopedStore()
	ctx := context.Background()

	if _, err := store.Set(ctx, "local:device-a", "alice", "first"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := store.Set(ctx, "local:device-b", "bob", "second"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	gotA, err := store.Get(ctx, "local:device-a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotA.Content != "first" {
		t.Fatalf("expected first content, got %s", gotA.Content)
	}

	gotB, err := store.Get(ctx, "local:device-b")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotB.Content != "second" {
		t.Fatalf("expected second content, got %s", gotB.Content)
	}

	if _, err := store.Get(ctx, "local:device-c"); err != ErrEmpty {
		t.Fatalf("expected ErrEmpty for unknown scope, got %v", err)
	}
}
