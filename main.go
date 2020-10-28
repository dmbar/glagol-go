package main

import (
	"fmt"
	"net/http"
)

const apiVersion = "v1"

func handleGETByOID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "inside Page route\n%q\n%q", r.Method, r.UserAgent())
}

func main() {
	http.HandleFunc(fmt.Sprintf("/api/%s/page", apiVersion), handleGETByOID)

	http.ListenAndServe(":80", nil)
}
