package solution

import (
	"context"
	"database/sql"
	"errors"
	"slices"
	"testing"
)

func newStore(t *testing.T) *UserStore {
	t.Helper()
	reset()
	db, err := sql.Open("fake", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	return NewUserStore(db)
}

func TestAddAndName(t *testing.T) {
	ctx := context.Background()
	s := newStore(t)

	id1, err := s.Add(ctx, "ada")
	id2, _ := s.Add(ctx, "bob")
	if err != nil || id1 != 1 || id2 != 2 {
		t.Fatalf("Add ids = %d, %d, err %v", id1, id2, err)
	}
	if got, err := s.Name(ctx, 2); err != nil || got != "bob" {
		t.Errorf("Name(2) = %q, %v", got, err)
	}
	if _, err := s.Name(ctx, 99); !errors.Is(err, ErrNotFound) {
		t.Errorf("Name(99) err = %v, want ErrNotFound", err)
	}
}

func TestList(t *testing.T) {
	ctx := context.Background()
	s := newStore(t)
	if got, err := s.List(ctx); err != nil || len(got) != 0 {
		t.Errorf("empty List() = %v, %v", got, err)
	}
	s.Add(ctx, "ada")
	s.Add(ctx, "bob")
	if got, _ := s.List(ctx); !slices.Equal(got, []string{"ada", "bob"}) {
		t.Errorf("List() = %v", got)
	}
}

func TestAddAllRollsBack(t *testing.T) {
	ctx := context.Background()
	s := newStore(t)

	if err := s.AddAll(ctx, []string{"x", "", "y"}); err == nil {
		t.Fatal("empty name should fail")
	}
	if got, _ := s.List(ctx); len(got) != 0 {
		t.Errorf("rollback failed, rows left: %v", got)
	}

	if err := s.AddAll(ctx, []string{"a", "b"}); err != nil {
		t.Fatal(err)
	}
	if got, _ := s.List(ctx); !slices.Equal(got, []string{"a", "b"}) {
		t.Errorf("List() = %v, want [a b]", got)
	}
}
