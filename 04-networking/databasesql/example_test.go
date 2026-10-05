package databasesqlref_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func open() *sql.DB {
	db, err := sql.Open("fake", "") // does not connect yet
	if err != nil {
		panic(err)
	}
	return db
}

const insert = "INSERT INTO users(name) VALUES(?)"

// Example_exec inserts a row and reads the result.
func Example_exec() {
	reset()
	db := open()
	defer db.Close()

	fmt.Println(db.Ping()) // the first real connection happens here

	res, err := db.Exec(insert, "ada")
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ := res.LastInsertId()
	n, _ := res.RowsAffected()
	fmt.Println(id, n)
	// Output:
	// <nil>
	// 1 1
}

// Example_query loops over several rows.
func Example_query() {
	reset()
	db := open()
	defer db.Close()
	for _, name := range []string{"ada", "bob", "cy"} {
		db.Exec(insert, name)
	}

	rows, err := db.Query("SELECT id, name FROM users ORDER BY id")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close() // releases the connection

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, name)
	}
	if err := rows.Err(); err != nil { // errors that ended the iteration early
		fmt.Println(err)
	}
	// Output:
	// 1 ada
	// 2 bob
	// 3 cy
}

// Example_queryRow reads exactly one row and handles "no rows".
func Example_queryRow() {
	reset()
	db := open()
	defer db.Close()
	db.Exec(insert, "ada")

	var name string
	err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
	fmt.Println(name, err)

	err = db.QueryRow("SELECT name FROM users WHERE id = ?", 99).Scan(&name)
	fmt.Println(errors.Is(err, sql.ErrNoRows), err)
	// Output:
	// ada <nil>
	// true sql: no rows in result set
}

// Example_prepared reuses one statement for many executions.
func Example_prepared() {
	reset()
	db := open()
	defer db.Close()

	stmt, err := db.Prepare(insert)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	for _, name := range []string{"ada", "bob"} {
		stmt.Exec(name)
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	fmt.Println(count)
	// Output: 2
}

// Example_transaction shows rollback and commit.
func Example_transaction() {
	reset()
	db := open()
	defer db.Close()

	count := func() (n int) {
		db.QueryRow("SELECT COUNT(*) FROM users").Scan(&n)
		return n
	}

	tx, _ := db.Begin()
	tx.Exec(insert, "temp")
	var inside int
	tx.QueryRow("SELECT COUNT(*) FROM users").Scan(&inside)
	tx.Rollback()
	fmt.Println(inside, count()) // seen inside the transaction, gone after rollback

	tx, _ = db.Begin()
	defer tx.Rollback() // safety net: no-op after Commit
	tx.Exec(insert, "keep")
	fmt.Println(tx.Commit(), count())
	// Output:
	// 1 0
	// <nil> 1
}

// Example_context shows that a cancelled context stops a call before it starts.
func Example_context() {
	db := open()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := db.ExecContext(ctx, insert, "x")
	fmt.Println(err)
	// Output: context canceled
}

// Example_pool configures and inspects the connection pool.
func Example_pool() {
	db := open()
	defer db.Close()

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(30 * time.Minute)

	fmt.Println(db.Stats().MaxOpenConnections)
	// Output: 5
}

// Example_nulls scans nullable values. These types work without a database.
func Example_nulls() {
	var s sql.NullString
	s.Scan(nil)
	fmt.Println(s.Valid, s.String == "")

	s.Scan("ada")
	fmt.Println(s.Valid, s.String)

	var n sql.Null[int] // generic form, Go 1.22+
	n.Scan(int64(7))
	fmt.Println(n.Valid, n.V)
	// Output:
	// false true
	// true ada
	// true 7
}
