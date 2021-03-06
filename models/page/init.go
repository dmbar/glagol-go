package page

import (
	"database/sql"
	"log"

	"github.com/dmbar/glagol-go/models"
)

// PageSourceDB is for working with Page objects in DB
var PageSourceDB models.DataSourceDB

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

	PageSourceDB = models.DataSourceDB{
		DB: models.DB,
		Stmnts: map[string]*sql.Stmt{
			"selectByOID": selectByOIDStmt,
			"insert":      insertStmt,
			"update":      updateStmt,
		},
	}
}
