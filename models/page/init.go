package page

import (
	"database/sql"
	"log"

	"github.com/dmbar/glagol-go/models"
)

// DBStorePage is for working with Page objects in DB
var DBStorePage models.DBStore

type crud struct {
	GetByOID func(oid string, store models.Store) (Page, error)
	Save     func(data Data, store models.Store) (Page, error)
	Update   func(oid string, data Data, store models.Store) (Page, error)
}

// CRUD object is used to invoke CRUD commands
var CRUD crud = crud{
	GetByOID: getByOID,
	Save:     save,
	Update:   update,
}

func init() {
	selectByOIDStmt, err := models.DB.Prepare(`SELECT created_on, updated_on, data FROM objects.pages WHERE oid = $1`)
	if err != nil {
		log.Fatal("prepare SELECT statement error:\n", err)
	}
	insertStmt, err := models.DB.Prepare(`INSERT INTO objects.pages(data) VALUES ($1) RETURNING oid, created_on, updated_on`)
	if err != nil {
		log.Fatal("prepare INSERT statement error:\n", err)
	}
	updateStmt, err := models.DB.Prepare(`UPDATE objects.pages SET data=$1 WHERE oid=$2 RETURNING created_on, updated_on`)
	if err != nil {
		log.Fatal("prepare UPDATE statement error:\n", err)
	}

	DBStorePage = models.DBStore{
		DB: models.DB,
		Stmnts: map[string]*sql.Stmt{
			"selectByOID": selectByOIDStmt,
			"insert":      insertStmt,
			"update":      updateStmt,
		},
	}
}
