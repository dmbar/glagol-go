package main

import (
	"fmt"
	"log"

	"github.com/bbneizvest/glagol/api/models/page"
)

func main() {
	// result, err := page.SelectByOID("dbc056a1-3115-4aba-b6ea-8f3269a34cbe", page.CRUD)
	// if err != nil {
	// 	log.Fatal("error SelectByOID", err)
	// }

	// fmt.Printf("result: %+v\n", result)

	// jsonData := []byte(`{
	// 	"header": "BROCKHAMPTON",
	// 	"body": "NO HALO"
	// }`)
	// result, err := page.Insert(jsonData, page.CRUD)
	// if err != nil {
	// 	log.Fatal("error Insert", err)
	// }

	// fmt.Printf("result: %+v\n", result)

	jsonData := []byte(`{
		"header": "Boards of Canada",
		"body": "Dayvan Cowboy"
	}`)
	result, err := page.Update("7f275b8d-9efa-46ee-8719-48cf6eff04d3", jsonData, page.CRUD)
	if err != nil {
		log.Fatal("error Update", err)
	}

	fmt.Printf("result: %+v\n", result)
}
