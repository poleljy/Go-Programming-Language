package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
)

type CentralDB struct {
	Conn *sql.DB
}

func InitDB(host, port, user, pwd, dbName string) (*CentralDB, error) {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pwd, host, port, dbName)
	db, err := sql.Open("pgx", conn)
	if err != nil {
		return nil, err
	}

	return &CentralDB{Conn: db}, nil
}

func (db CentralDB) Disconnect() {
	if err := db.Conn.Close(); err != nil {
		log.Println("Failed to disconnect db:", err)
	}
}

func (db CentralDB) IsAvailable() bool {
	if err := db.Conn.Ping(); err != nil {
		return false
	}
	return true
}

func (db CentralDB) Query() {
	var sum, n int32

	// invoke query
	rows, err := db.Conn.Query("SELECT generate_series(1,$1)", 10)
	// handle query error
	if err != nil {
		fmt.Println(err)
		if pgErr, ok := err.(pgx.PgError); ok {
			if pgErr.Code == "0A000" {
				return
			}
		}
	}
	// defer close result set
	defer rows.Close()

	// Iter results
	for rows.Next() {
		if err = rows.Scan(&n); err != nil {
			fmt.Println(err) // Handle scan error
		}
		sum += n // Use result
	}

	// check iteration error
	if rows.Err() != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func (db CentralDB) QueryRow() {
	var sum int
	err := db.Conn.QueryRow(`SELECT sum(n) FROM (SELECT generate_series(1,$1) as n) a;`, 10).Scan(&sum)
	//err := db.Conn.QueryRow(`select age from test where name = $1`, "Lee").Scan(&sum)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("NO data return:", err)
			return
		}
		fmt.Println(err)
	}
	fmt.Println(sum)
}

func main() {
	db, err := InitDB("123.59.135.227", "5432", "postgres", "otitan123", "TitanCloud.System")
	if err != nil {
		log.Fatal("Failed to initialize central db :", err)
	}

	defer db.Disconnect()

	if db.IsAvailable() {
		fmt.Println("Connection is available")
	} else {
		fmt.Println("Connection is unavailable")
	}

	db.Query()

	db.QueryRow()
}
