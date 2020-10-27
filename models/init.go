package models

import (
	"database/sql"
	"log"

	// Loading PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB is a sql.DB object used for DB connections
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("pgx", "postgres://api_user:api_user@localhost:5432/glagol")
	if err != nil {
		log.Fatal("DB connection error:\n", err)
	}
}

// CRUDForObject is an abstraction layer for common CRUD operations
// Each object has it's own implementation in it's own package
//
// Also simplifies mocking :)
type CRUDForObject interface {
	Select(params []interface{}, dest ...interface{}) error
	Insert(params []interface{}, dest ...interface{}) error
	Update(params []interface{}, dest ...interface{}) error
}
