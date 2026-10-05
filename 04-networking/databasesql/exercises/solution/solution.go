// Package solution holds the reference solutions for the database/sql exercises.
package solution

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ErrNotFound is returned when a user does not exist.
var ErrNotFound = errors.New("user not found")

// UserStore stores user names in the users table.
type UserStore struct{ db *sql.DB }

// NewUserStore wraps db.
func NewUserStore(db *sql.DB) *UserStore { return &UserStore{db: db} }

// Add inserts a user and returns its new id.
func (s *UserStore) Add(ctx context.Context, name string) (int64, error) {
	res, err := s.db.ExecContext(ctx, "INSERT INTO users(name) VALUES(?)", name)
	if err != nil {
		return 0, fmt.Errorf("add user: %w", err)
	}
	return res.LastInsertId()
}

// Name returns the name of user id, or ErrNotFound.
func (s *UserStore) Name(ctx context.Context, id int64) (string, error) {
	var name string
	err := s.db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", id).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("get user %d: %w", id, err)
	}
	return name, nil
}

// List returns all names ordered by id.
func (s *UserStore) List(ctx context.Context) ([]string, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name FROM users ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, rows.Err()
}

// AddAll inserts all names in one transaction. If any name is empty, nothing
// is inserted.
func (s *UserStore) AddAll(ctx context.Context, names []string) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // no-op after a successful Commit

	for _, name := range names {
		if name == "" {
			return errors.New("add all: empty name")
		}
		if _, err := tx.ExecContext(ctx, "INSERT INTO users(name) VALUES(?)", name); err != nil {
			return err
		}
	}
	return tx.Commit()
}
