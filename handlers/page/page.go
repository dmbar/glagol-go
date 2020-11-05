package page

import (
	"fmt"
	"net/http"
)

// HandleGET handles all operations for GET method
func HandleGET(w http.ResponseWriter, rq *http.Request) {
	oid := rq.URL.Query().Get("oid")

	if oid == "" {
		fmt.Fprintf(w, "OID was not provided or it is not a valid UUID value")
		return
	}

	fmt.Fprintf(w, "good morning slut")
}
