// Package page provides functions to interact with storage
// and perform CRUD operations on Page object
package page

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dmbar/glagol-go/models"
)

const typeAssertError string = "query page error - type assertion failed at %v"

// Page represents one (1) Page
type Page struct {
	Meta meta `json:"meta"`
	Data Data `json:"data"`
}

type meta struct {
	OID       string    `json:"oid"`
	CreatedOn time.Time `json:"createdOn"`
	UpdatedOn time.Time `json:"updatedOn"`
}

// Data represents contents of Page object. Corresponds to "Data" column in DB
type Data struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

// GetByOID returns one (1) Page from source using ObjectID
//  oid - ObjectID of record
//  store - data source (DB, external service, etc.)
func GetByOID(oid string, store models.Store) (Page, error) {
	var result Page
	// Variables to be used in query
	var createdOn time.Time
	var updatedOn time.Time
	var dataJSON []byte

	// Executing query to DB
	queryResult, err := store.ExecuteQuery("selectByOID", []interface{}{oid}, &createdOn, &updatedOn, &dataJSON)
	if err != nil {
		return Page{}, err
	}

	result.Meta.OID = oid
	// Parsing query result
	// Since UUIDs are unique, queryResult will be an array with 1 element
	row, ok := queryResult[0].([]interface{})
	if !ok {
		return Page{}, fmt.Errorf(typeAssertError, row)
	}
	for colNum, column := range row {
		switch colNum {
		case 0: // CreatedOn
			value, ok := column.(*time.Time)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "CreatedOn")
			}
			result.Meta.CreatedOn = *value
		case 1: // UpdatedOn
			value, ok := column.(*time.Time)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "UpdatedOn")
			}
			result.Meta.UpdatedOn = *value
		case 2: // JSON Data
			value, ok := column.(*[]byte)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "Data")
			}
			dataJSON = *value
		}
	}

	// JSON Unmarshaling
	err = json.Unmarshal(dataJSON, &result.Data)
	if err != nil {
		return Page{}, fmt.Errorf("JSON unmarshaling error: %v", err)
	}

	return result, nil
}

// Save saves Page data to source and returns new Page
//  data - Page data
//  store - data source (DB, external service, etc.)
func Save(data Data, store models.Store) (Page, error) {
	var result Page
	// Variables to be used in query
	var oid string
	var createdOn time.Time
	var updatedOn time.Time

	// Executing query to DB
	queryResult, err := store.ExecuteQuery("insert", []interface{}{data}, &oid, &createdOn, &updatedOn)
	if err != nil {
		return Page{}, err
	}

	result.Data = data
	// Parsing query result
	// Since we are saving 1 instance of Page, queryResult will be an array with 1 element
	row, ok := queryResult[0].([]interface{})
	if !ok {
		return Page{}, fmt.Errorf(typeAssertError, row)
	}
	for colNum, column := range row {
		switch colNum {
		case 0: // JSON Data
			value, ok := column.(*string)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "OID")
			}
			result.Meta.OID = *value
		case 1: // CreatedOn
			value, ok := column.(*time.Time)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "CreatedOn")
			}
			result.Meta.CreatedOn = *value
		case 2: // UpdatedOn
			value, ok := column.(*time.Time)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "UpdatedOn")
			}
			result.Meta.UpdatedOn = *value
		}
	}

	return result, nil
}

// Update updates Page instance in source and returns updated Page
//  oid - ObjectID of record
//  data - Page data
//  store - data source (DB, external service, etc.)
func Update(oid string, data Data, store models.Store) (Page, error) {
	var result Page
	// Variables to be used in query
	var createdOn time.Time
	var updatedOn time.Time

	// Executing query to DB
	queryResult, err := store.ExecuteQuery("update", []interface{}{data, oid}, &createdOn, &updatedOn)
	if err != nil {
		return Page{}, err
	}

	result.Meta.OID = oid
	result.Data = data
	// Since UUIDs are unique, queryResult will be an array with 1 element
	row, ok := queryResult[0].([]interface{})
	if !ok {
		return Page{}, fmt.Errorf(typeAssertError, row)
	}
	for colNum, column := range row {
		switch colNum {
		case 0: // CreatedOn
			value, ok := column.(*time.Time)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "CreatedOn")
			}
			result.Meta.CreatedOn = *value
		case 1: // UpdatedOn
			value, ok := column.(*time.Time)
			if !ok {
				return Page{}, fmt.Errorf(typeAssertError, "UpdatedOn")
			}
			result.Meta.UpdatedOn = *value
		}
	}

	return result, nil
}
