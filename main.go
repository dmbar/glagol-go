package main

import (
	"fmt"
	"net/http"

	"github.com/dmbar/glagol-go/handlers/page"
)

const apiRoot string = "/api"
const apiVersion string = "/v1"

func main() {
	http.HandleFunc(fmt.Sprintf("%s%s/page", apiRoot, apiVersion), page.HandleGET)

	http.ListenAndServe("localhost:8080", nil)
}
