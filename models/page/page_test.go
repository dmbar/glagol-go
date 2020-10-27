package page

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

// * Constants *
const (
	tOid             string = "2aec0be7-b093-44eb-8b6d-9367d2ef7a5b"
	tHeader          string = "Eraserhead"
	tBody            string = "On 19 March 1977, the world changed, after which there was a long uncomfortable silence."
	tGenericErrorMsg string = "jUsT aN eRrrAwr :3"
)

var tPage = Page{
	Meta: meta{
		Oid:       tOid,
		CreatedOn: time.Date(1994, 2, 1, 13, 13, 13, 0, time.UTC),
		UpdatedOn: time.Date(1994, 2, 1, 13, 13, 13, 0, time.UTC),
	},
	Data: data{
		Header: tHeader,
		Body:   tBody,
	},
}
var tJSONData = []byte(fmt.Sprintf(`{"header": "%s", "body": "%s"}`, tHeader, tBody))

// * TESTS *
func TestSelectByOID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCRUD := mockCRUDObject{scenario: Success}

		got, err := SelectByOID(tOid, mockCRUD)
		want := tPage

		if err != nil {
			t.Errorf("expected success, got error %v\n", err)
		}

		if got != want {
			t.Errorf("\ngot:\n%+v\n\nwant:\n%+v", got, want)
		}
	})

	t.Run("page not found", func(t *testing.T) {
		mockCRUD := mockCRUDObject{scenario: NoRowsError}

		got, err := SelectByOID("2aec0be7-b093-44eb-8b6d-9367d2ef7a5b", mockCRUD)
		want := Page{}

		if got != want && err != sql.ErrNoRows {
			t.Errorf("expected empty Page and error ErrNoRows, got Page: %+v; error: %v\n", got, err)
		}
	})

	t.Run("generic error", func(t *testing.T) {
		mockCRUD := mockCRUDObject{scenario: Error}

		got, err := SelectByOID("2aec0be7-b093-44eb-8b6d-9367d2ef7a5b", mockCRUD)
		want := Page{}

		if got != want && err.Error() != tGenericErrorMsg {
			t.Errorf("expected empty Page and error, got Page: %+v; error: %v\n", got, err)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCRUD := mockCRUDObject{scenario: Success}

		got, err := Insert(tJSONData, mockCRUD)
		want := tPage

		if err != nil {
			t.Errorf("expected success, got error %v\n", err)
		}

		if got != want {
			t.Errorf("\ngot:\n%+v\n\nwant:\n%+v", got, want)
		}
	})

	t.Run("generic error", func(t *testing.T) {
		mockCRUD := mockCRUDObject{scenario: Error}

		got, err := Insert(tJSONData, mockCRUD)
		want := Page{}

		if got != want && err.Error() != tGenericErrorMsg {
			t.Errorf("expected empty Page and error, got Page: %+v; error: %v\n", got, err)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCRUD := mockCRUDObject{scenario: Success}

		got, err := Update(tOid, tJSONData, mockCRUD)
		want := tPage

		if err != nil {
			t.Errorf("expected success, got error %v\n", err)
		}

		if got != want {
			t.Errorf("\ngot:\n%+v\n\nwant:\n%+v", got, want)
		}
	})

	t.Run("generic error", func(t *testing.T) {
		mockCRUD := mockCRUDObject{scenario: Error}

		got, err := Update(tOid, tJSONData, mockCRUD)
		want := Page{}

		if got != want && err.Error() != tGenericErrorMsg {
			t.Errorf("expected empty Page and error, got Page: %+v; error: %v\n", got, err)
		}
	})
}

func TestQueryStrings(t *testing.T) {
	gotList := []string{selectByOIDQueryString, insertQueryString, updateQueryString}
	want := []string{"SELECT oid, created_on, updated_on, data FROM objects.pages WHERE oid = $1",
		"INSERT INTO objects.pages(data) VALUES ($1) RETURNING oid, created_on, updated_on, data",
		"UPDATE objects.pages SET data=$1 WHERE oid=$2 RETURNING oid, created_on, updated_on, data"}

	for idx, got := range gotList {
		if got != want[idx] {
			t.Errorf("\ngot:\n%s\nwant:\n%s", got, want[idx])
		}
	}
}

// * Mocking *
// scenarios
const (
	Success = iota
	NoRowsError
	Error
)

type mockCRUDObject struct {
	scenario int
}

func (mock mockCRUDObject) Select(params []interface{}, dest ...interface{}) error {
	switch mock.scenario {
	case Success:
		stubScanValues(dest...)
	case NoRowsError:
		return sql.ErrNoRows
	case Error:
		return fmt.Errorf(tGenericErrorMsg)
	}

	return nil
}

func (mock mockCRUDObject) Insert(params []interface{}, dest ...interface{}) error {
	switch mock.scenario {
	case Success:
		stubScanValues(dest...)
	case Error:
		return fmt.Errorf(tGenericErrorMsg)
	}

	return nil
}

func (mock mockCRUDObject) Update(params []interface{}, dest ...interface{}) error {
	switch mock.scenario {
	case Success:
		stubScanValues(dest...)
	case Error:
		return fmt.Errorf(tGenericErrorMsg)
	}

	return nil
}

// Stub Scan values
func stubScanValues(dest ...interface{}) error {
	for idx, val := range dest {
		switch idx {
		case 0: // * Oid
			currVal, ok := val.(*string)
			if !ok {
				return fmt.Errorf("error in stub - expected type *string, got %T", val)
			}
			*currVal = tOid
		case 1: // * CreatedOn
			currVal, ok := val.(*time.Time)
			if !ok {
				return fmt.Errorf("error in stub - expected type *time.Time, got %T", val)
			}
			*currVal = time.Date(1994, 2, 1, 13, 13, 13, 0, time.UTC)
		case 2: // * UpdatedOn
			currVal, ok := val.(*time.Time)
			if !ok {
				return fmt.Errorf("error in stub - expected type *time.Time, got %T", val)
			}
			*currVal = time.Date(1994, 2, 1, 13, 13, 13, 0, time.UTC)
		case 3: // * Data
			currVal, ok := val.(*[]byte)
			if !ok {
				return fmt.Errorf("error in stub - expected type *[]bytes, got %T", val)
			}
			*currVal = []byte(fmt.Sprintf(`{
				"header": "%s",
				"body":   "%s"
				}`, tHeader, tBody))
		}
	}

	return nil
}
