package databasesqlref_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"strings"
	"sync"
)

// This file is a tiny in-memory database driver so the examples run without a
// real server. It has one table, users(id, name), and recognizes only the
// statements used in the examples. It is NOT a SQL engine. In real code you
// import a driver such as github.com/jackc/pgx/v5/stdlib instead.

type fakeDB struct {
	mu    sync.Mutex
	names []string // id = index + 1
}

var fake = &fakeDB{}

func reset() {
	fake.mu.Lock()
	fake.names = nil
	fake.mu.Unlock()
}

func init() { sql.Register("fake", fakeDriver{}) }

type fakeDriver struct{}

func (fakeDriver) Open(string) (driver.Conn, error) { return &fakeConn{}, nil }

type fakeConn struct{ snapshot []string }

func (c *fakeConn) Prepare(q string) (driver.Stmt, error) {
	return &fakeStmt{q: strings.TrimSpace(q)}, nil
}
func (c *fakeConn) Close() error { return nil }
func (c *fakeConn) Begin() (driver.Tx, error) {
	fake.mu.Lock()
	defer fake.mu.Unlock()
	c.snapshot = append([]string(nil), fake.names...)
	return &fakeTx{c}, nil
}

type fakeTx struct{ c *fakeConn }

func (t *fakeTx) Commit() error { return nil }
func (t *fakeTx) Rollback() error {
	fake.mu.Lock()
	fake.names = t.c.snapshot
	fake.mu.Unlock()
	return nil
}

type fakeStmt struct{ q string }

func (s *fakeStmt) Close() error  { return nil }
func (s *fakeStmt) NumInput() int { return strings.Count(s.q, "?") }

type fakeResult struct{ id, n int64 }

func (r fakeResult) LastInsertId() (int64, error) { return r.id, nil }
func (r fakeResult) RowsAffected() (int64, error) { return r.n, nil }

func (s *fakeStmt) Exec(args []driver.Value) (driver.Result, error) {
	fake.mu.Lock()
	defer fake.mu.Unlock()
	switch {
	case strings.HasPrefix(s.q, "INSERT INTO users"):
		fake.names = append(fake.names, args[0].(string))
		return fakeResult{int64(len(fake.names)), 1}, nil
	case strings.HasPrefix(s.q, "DELETE FROM users"):
		n := int64(len(fake.names))
		fake.names = nil
		return fakeResult{0, n}, nil
	}
	return nil, errors.New("fake driver: unsupported statement: " + s.q)
}

type fakeRows struct {
	cols []string
	data [][]driver.Value
	i    int
}

func (r *fakeRows) Columns() []string { return r.cols }
func (r *fakeRows) Close() error      { return nil }
func (r *fakeRows) Next(dest []driver.Value) error {
	if r.i >= len(r.data) {
		return io.EOF
	}
	copy(dest, r.data[r.i])
	r.i++
	return nil
}

func (s *fakeStmt) Query(args []driver.Value) (driver.Rows, error) {
	fake.mu.Lock()
	defer fake.mu.Unlock()
	switch {
	case strings.HasPrefix(s.q, "SELECT id, name FROM users"):
		rows := &fakeRows{cols: []string{"id", "name"}}
		for i, n := range fake.names {
			rows.data = append(rows.data, []driver.Value{int64(i + 1), n})
		}
		return rows, nil
	case strings.HasPrefix(s.q, "SELECT name FROM users WHERE id"):
		rows := &fakeRows{cols: []string{"name"}}
		id := args[0].(int64)
		if id >= 1 && int(id) <= len(fake.names) {
			rows.data = append(rows.data, []driver.Value{fake.names[id-1]})
		}
		return rows, nil
	case strings.HasPrefix(s.q, "SELECT COUNT(*) FROM users"):
		return &fakeRows{cols: []string{"count"}, data: [][]driver.Value{{int64(len(fake.names))}}}, nil
	}
	return nil, errors.New("fake driver: unsupported query: " + s.q)
}
