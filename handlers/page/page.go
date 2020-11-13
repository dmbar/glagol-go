package page

import (
	"fmt"
	"net/http"

	"github.com/dmbar/glagol-go/models/page"
	"github.com/google/uuid"
)

const invalidOID = "OID was not provided or it is not a valid UUID value"
const internalServerError = "Internal server error"

// HandleGET handles all operations for GET method
func HandleGET(w http.ResponseWriter, rq *http.Request) {
	oid := rq.URL.Query().Get("oid")

	// Validating OID
	if _, err := uuid.Parse(oid); err != nil {
		fmt.Fprintf(w, invalidOID)
		return
	}

	_, err := page.CRUD.GetByOID(oid, page.DBStorePage)
	if err != nil {
		fmt.Fprintf(w, internalServerError)
		return
	}
}
