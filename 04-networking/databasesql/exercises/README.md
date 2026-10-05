# database/sql exercises

Try each exercise yourself first, then compare with
[solution/solution.go](solution/solution.go). Check the reference solutions with:

```bash
go test ./04-networking/databasesql/exercises/...
```

## Exercise: UserStore

Implement a `UserStore` on top of `*sql.DB`:

```go
func NewUserStore(db *sql.DB) *UserStore
func (s *UserStore) Add(ctx context.Context, name string) (int64, error)
func (s *UserStore) Name(ctx context.Context, id int64) (string, error)
func (s *UserStore) List(ctx context.Context) ([]string, error)
func (s *UserStore) AddAll(ctx context.Context, names []string) error
```

SQL to use (the test driver recognizes exactly these):

```sql
INSERT INTO users(name) VALUES(?)
SELECT name FROM users WHERE id = ?
SELECT id, name FROM users ORDER BY id
```

Rules:

- `Add` returns the new id (`LastInsertId`)
- `Name` returns the exported error `ErrNotFound` when the id does not exist
  (`sql.ErrNoRows`), and wraps other errors with context
- `List` closes the rows and checks `rows.Err()`
- `AddAll` inserts every name **in one transaction**; if any name is empty it
  returns an error and **nothing** is inserted
- Use the `...Context` methods everywhere

Hints: `defer tx.Rollback()` right after `BeginTx`, `errors.Is(err, sql.ErrNoRows)`.
The in-memory driver is in `fakedriver_test.go`.

## Bonus
Run the same code against a real PostgreSQL database (Docker `postgres:16`)
using `pgx`'s `database/sql` adapter. What changes? (Placeholders, `LastInsertId`.)
