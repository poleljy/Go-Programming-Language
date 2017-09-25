package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx"
)

func InitDbConfig(dbHost string, dbPort int, dbUser, dbPassword, dbDatabase string) pgx.ConnConfig {
	config := pgx.ConnConfig{
		Host:     dbHost,
		Port:     uint16(dbPort),
		Database: dbDatabase,
		User:     dbUser,
		Password: dbPassword,
	}
	return config
}

// pgx
type PgxUtil struct {
	Config pgx.ConnConfig
}

func GetFieldIndex(rows *pgx.Rows, name string) int {
	fds := rows.FieldDescriptions()
	for i := 0; i < len(fds); i++ {
		if strings.ToUpper(fds[i].Name) == strings.ToUpper(name) {
			return i
		}
	}
	return -1
}

func (pgxdb *PgxUtil) Query(sql string) (*pgx.Conn, *pgx.Rows, error) {
	// Connect
	conn, err := pgx.Connect(pgxdb.Config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		return nil, nil, err
	}

	// Query
	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Printf("Failed to query(%s): %v\n", sql, err)
		return conn, nil, err
	}
	return conn, rows, nil
}

func (pgxdb *PgxUtil) Exec(sql string) (int64, error) {
	var rowsAffected int64 = 0

	// Connect
	conn, err := pgx.Connect(pgxdb.Config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		return rowsAffected, err
	}
	defer conn.Close()

	// Exec
	commandTag, err := conn.Exec(sql)
	if err != nil {
		fmt.Printf("Failed to exec sql(%s): %v\n", sql, err)
		return rowsAffected, err
	}
	return commandTag.RowsAffected(), nil
}

func Tutorial_pgx() {
	config := InitDbConfig("192.168.1.189", 5432, "postgres", "123456", "TitanSBO")
	db := PgxUtil{Config: config}

	sql := `insert into groups(group_name,address,country,join_time) 
						values('AeroSpaceTitan','研发部门','China',current_timestamp) RETURNING group_id`
	db.Exec(sql)
}
