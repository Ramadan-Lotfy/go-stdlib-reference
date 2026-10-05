# database/sql

## Overview
`database/sql` is Go's database access layer. It defines a generic API (`DB`,
`Tx`, `Rows`, `Stmt`) and a connection pool; a separate **driver** package
(imported for its side effect) talks to the actual database. `sql.DB` is not a
single connection: it is a pool, safe for concurrent use, created once and kept
for the life of the program.

The examples here run against a tiny in-memory driver defined in
`fakedriver_test.go`, so no database server is needed. With a real PostgreSQL
database only the import and the connection string change:

```go
import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // registers the driver named "pgx"
)

db, err := sql.Open("pgx", "postgres://user:pass@localhost:5432/app?sslmode=disable")
if err != nil { /* bad DSN */ }
defer db.Close()
if err := db.PingContext(ctx); err != nil { /* cannot connect */ }

rows, err := db.QueryContext(ctx, "SELECT id, name FROM users WHERE age > $1", 18)
```

PostgreSQL uses `$1, $2` placeholders; MySQL and SQLite use `?`.

## Core API

| Function / Type | Purpose |
|---|---|
| `sql.Open(driver, dsn)` | Create the pool (does **not** connect); `db.Ping()` to check |
| `db.Exec`, `ExecContext` | INSERT/UPDATE/DELETE: returns `Result` (`LastInsertId`, `RowsAffected`) |
| `db.Query`, `QueryContext` | Multiple rows: loop `rows.Next()` + `rows.Scan(...)`; `defer rows.Close()`; check `rows.Err()` |
| `db.QueryRow(...).Scan(...)` | Exactly one row; returns `sql.ErrNoRows` if none |
| `db.Prepare` | Reusable statement |
| `db.Begin`, `BeginTx`, `tx.Commit`, `tx.Rollback` | Transactions |
| `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime` | Pool tuning |
| `sql.NullString`, `NullInt64`, `sql.Null[T]` | Nullable columns |
| `db.Stats()` | Pool statistics |

## Examples
See [example_test.go](example_test.go). Run them with:

```bash
go test -v ./04-networking/databasesql/
```

## Gotchas
- **Always use placeholders** (`$1` or `?`) for values. Never build SQL with `fmt.Sprintf` or `+`: that is SQL injection. (Table and column names cannot be placeholders; validate them against an allow-list.)
- **`sql.Open` does not connect.** A wrong password only shows up on the first query or `Ping`.
- **Always `defer rows.Close()`** right after checking the error, and **check `rows.Err()`** after the loop. An unclosed `Rows` holds a pool connection forever.
- **`QueryRow` defers its error to `Scan`.** Test `errors.Is(err, sql.ErrNoRows)` to detect a missing row; do not treat it as a failure by default.
- **Use the `...Context` variants** so a cancelled request or a timeout releases the query.
- **Always pair `Begin` with `defer tx.Rollback()`.** After a successful `Commit` the rollback is a harmless no-op; on any early return it prevents a leaked transaction.
- **Inside a transaction use `tx`, not `db`.** Calls on `db` run on a different connection and do not see the transaction.
- **Nullable columns need `sql.Null*` types or pointer destinations,** or `Scan` fails when the value is NULL.
- **Tune the pool:** the default is unlimited open connections, which can overwhelm the database. Set `SetMaxOpenConns` (and a `ConnMaxLifetime`) based on your server's limits.
- **`LastInsertId` is not supported by every driver** (PostgreSQL: use `INSERT ... RETURNING id` with `QueryRow`).
- **Create one `*sql.DB` per database,** share it, close it on shutdown. Do not open one per request.

## Java comparison

| Java | Go |
|---|---|
| `DataSource` / HikariCP | `*sql.DB` (pool built in) |
| `Connection.prepareStatement` | `db.Prepare` / placeholders in `Query` |
| `ResultSet.next()` + `getString` | `rows.Next()` + `rows.Scan(&s)` |
| try-with-resources | `defer rows.Close()` |
| `@Transactional` | `tx, _ := db.BeginTx(...)` + `defer tx.Rollback()` + `tx.Commit()` |
| `JdbcTemplate` / JPA | not in stdlib: see `sqlx`, `pgx`, `sqlc`, `GORM` |

## Exercises
See [exercises/](exercises).

## Links
- Docs: https://pkg.go.dev/database/sql
- Source: https://cs.opensource.google/go/go/+/master:src/database/sql/
