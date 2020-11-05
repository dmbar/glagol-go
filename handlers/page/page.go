package page

import (
	"fmt"
	"net/http"
)

func HandleGET(w http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(w, "=== Page GET ===\nRequest headers:\n\n")
	for name, headers := range rq.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%q: %q\n", name, h)
		}
	}

	oid := rq.URL.Query().Get("oid")
	fmt.Fprintf(w, "\n=== Query ===\nOID: %q", oid)
}
