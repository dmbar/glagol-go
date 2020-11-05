package page

import (
	"fmt"
	"net/http"
)

// HandleGET handles all operations for GET method
func HandleGET(w http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(w, "good morning slut")

}
