package page

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/bbneizvest/glagol/api/models"
)

type pageCRUD struct{}

// CRUD operations for Page
var CRUD pageCRUD

// Select executes SELECT statement to get single instance of Page
//
// params - Parameters to be passed to statement
//
// dest - Variables into which results should be saved
func (c pageCRUD) Select(params []interface{}, dest ...interface{}) error {
	return selectStmt.QueryRow(params...).Scan(dest...)
}

// Insert executes INSERT INTO statement to save and
// return new Page instance
//
// params - Parameters to be passed to statement
//
// dest - Variables into which results should be saved
func (c pageCRUD) Insert(params []interface{}, dest ...interface{}) error {
	return insertStmt.QueryRow(params...).Scan(dest...)
}

// Update executes UPDATE statement to save to and
// return updated Page instance
//
// params - Parameters to be passed to statement
//
// dest - Variables into which results should be saved
func (c pageCRUD) Update(params []interface{}, dest ...interface{}) error {
	return updateStmt.QueryRow(params...).Scan(dest...)
}

// SELECT statement to query Page instance
var selectStmt *sql.Stmt

// INSERT statement to save new Page instance
var insertStmt *sql.Stmt

// UPDATE statement to save new Page instance
var updateStmt *sql.Stmt

// Preparing SQL statements for execution
func init() {
	var err error

	selectStmt, err = models.DB.Prepare(`SELECT oid, created_on, updated_on, data FROM objects.pages WHERE oid = $1`)
	if err != nil {
		log.Fatal("prepare SELECT statement error:\n", err)
	}

	insertStmt, err = models.DB.Prepare(`INSERT INTO objects.pages(data) VALUES ($1) RETURNING oid, created_on, updated_on, data`)
	if err != nil {
		log.Fatal("prepare INSERT INTO statement error:\n", err)
	}

	updateStmt, err = models.DB.Prepare(`UPDATE objects.pages SET data=$1 WHERE oid=$2 RETURNING oid, created_on, updated_on, data`)
	if err != nil {
		log.Fatal("prepare UPDATE statement error:\n", err)
	}
}

// SelectByOID retrieves a single Page record from database.
// If no results are found then returns sql.ErrNoRows.
//
// oid - UUID key of a record
//
// crud - interface for sending queries
func SelectByOID(oid string, crud models.CRUDForObject) (Page, error) {
	var resultPage Page
	var jsonPageData []byte

	err := crud.Select([]interface{}{oid},
		&resultPage.Meta.Oid,
		&resultPage.Meta.CreatedOn,
		&resultPage.Meta.UpdatedOn,
		&jsonPageData)

	if err != nil {
		log.Printf("select query error: %v\n", err)
		return Page{}, err
	}

	err = json.Unmarshal(jsonPageData, &resultPage.Data)
	if err != nil {
		log.Printf("JSON unmarshaling error: %v\n", err)
		return Page{}, err
	}

	return resultPage, nil
}

// Insert saves and returns new Page record from DB.
//
// jsonData - Page Data field in JSON encoding
//
// crud - interface for sending queries
func Insert(jsonPageData []byte, crud models.CRUDForObject) (Page, error) {
	var resultPage Page

	err := crud.Insert([]interface{}{jsonPageData},
		&resultPage.Meta.Oid,
		&resultPage.Meta.CreatedOn,
		&resultPage.Meta.UpdatedOn,
		&jsonPageData)

	if err != nil {
		log.Printf("insert query error: %v\n", err)
		return Page{}, err
	}

	err = json.Unmarshal(jsonPageData, &resultPage.Data)
	if err != nil {
		log.Printf("JSON unmarshaling error: %v\n", err)
		return Page{}, err
	}

	return resultPage, nil
}

// Update saves page data to existing record
// and returns updated Page record from DB.
//
// oid - UUID key of a record
//
// jsonData - Page Data field in JSON encoding
//
// crud - interface for sending queries
func Update(oid string, jsonPageData []byte, crud models.CRUDForObject) (Page, error) {
	var resultPage Page

	err := crud.Update([]interface{}{jsonPageData, oid},
		&resultPage.Meta.Oid,
		&resultPage.Meta.CreatedOn,
		&resultPage.Meta.UpdatedOn,
		&jsonPageData)

	if err != nil {
		log.Printf("update query error: %v\n", err)
		return Page{}, err
	}

	err = json.Unmarshal(jsonPageData, &resultPage.Data)
	if err != nil {
		log.Printf("JSON unmarshaling error: %v\n", err)
		return Page{}, err
	}

	return resultPage, nil
}

// Page is an object for a single article with Header and Body.
type Page struct {
	Meta meta
	Data data
}
type meta struct {
	Oid       string
	CreatedOn time.Time
	UpdatedOn time.Time
}

type data struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}
