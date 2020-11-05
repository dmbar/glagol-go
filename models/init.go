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
		log.Fatal("database connection error:\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("ping database error:\n", err)
	}
}

// Store is a general interface for all queries to different data sources
type Store interface {
	// ExecuteQuery performs a query and returns rows with result.
	//
	// queryCode - Query to execute. Optional
	//
	// params - Query parameters
	//
	// dest - Variables for saving. Optional
	ExecuteQuery(queryCode string, params []interface{}, dest ...interface{}) ([]interface{}, error)
}

// DBStore is used for executing queries in database
type DBStore struct {
	DB     *sql.DB
	Stmnts map[string]*sql.Stmt
}

// ExecuteQuery performs a query and returns rows with result.
// If query is not supposed to return anything it returns (nil, nil)
//
// queryCode - Query to execute. Should be defined on DataSourceDB object
//
// params - Parameters for a SQL statement
//
// dest - Variables for saving results of a query, equals number of columns in SQL statement
func (q DBStore) ExecuteQuery(queryCode string, params []interface{}, dest ...interface{}) ([]interface{}, error) {
	rows, err := q.Stmnts[queryCode].Query(params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []interface{}
	for rows.Next() {
		if err = rows.Scan(dest...); err != nil {
			return nil, err
		}
		result = append(result, dest)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
