package page

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const invalidOID = "OID was not provided or it is not a valid UUID value"

// HandleGET handles all operations for GET method
func HandleGET(w http.ResponseWriter, rq *http.Request) {
	oid := rq.URL.Query().Get("oid")

	// Validating OID
	if _, err := uuid.Parse(oid); err != nil {
		fmt.Fprintf(w, invalidOID)
		return
	}

	fmt.Fprintf(w, "good morning slut") // Call to GetPageByOID
}
