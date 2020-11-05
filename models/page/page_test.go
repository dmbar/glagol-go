package page

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

const (
	SuccessGet = iota
	SuccessSave
	SuccessUpdate
	Error
	ErrorWrongType
)

var (
	tOID          = "413f2dcc-a360-450c-acd8-6588e2861e5b"
	tHeader       = "Eraserhead"
	tBody         = "On 19 March 1977, the world changed, after which there was a long uncomfortable silence."
	tErrorMsg     = "jUsT aN eRrrAwr tee-hee :3"
	tJSONErrorMsg = "JSON unmarshaling error: "
)

var tDataJSON = []byte(fmt.Sprintf(`{"header":%q,"body":%q}`, tHeader, tBody))
var tDataInvalidJSON = []byte("invalid json")

var tPage = Page{
	Meta: meta{
		OID:       tOID,
		CreatedOn: time.Date(1994, 2, 1, 13, 13, 13, 0, time.UTC),
		UpdatedOn: time.Date(1994, 2, 1, 13, 13, 13, 0, time.UTC),
	},
	Data: Data{
		Header: tHeader,
		Body:   tBody,
	},
}

type mockStore struct {
	scenario  int
	oid       interface{}
	createdOn interface{}
	updatedOn interface{}
	dataJSON  interface{}
}

func (q mockStore) ExecuteQuery(queryCode string, params []interface{}, dest ...interface{}) ([]interface{}, error) {
	switch q.scenario {
	case SuccessGet:
		row := []interface{}{q.createdOn, q.updatedOn, q.dataJSON}
		return []interface{}{row}, nil
	case SuccessSave:
		row := []interface{}{q.oid, q.createdOn, q.updatedOn}
		return []interface{}{row}, nil
	case SuccessUpdate:
		row := []interface{}{q.createdOn, q.updatedOn}
		return []interface{}{row}, nil
	case Error:
		return nil, fmt.Errorf(tErrorMsg)
	}

	return nil, fmt.Errorf("unexpected end for mock")
}

func TestGetByOID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := mockStore{
			scenario:  SuccessGet,
			createdOn: &tPage.Meta.CreatedOn,
			updatedOn: &tPage.Meta.UpdatedOn,
			dataJSON:  &tDataJSON,
		}
		want := tPage

		got, err := GetByOID(tOID, mock)

		if err != nil {
			t.Errorf("expected Page, got err: %q", err)
			t.FailNow()
		}

		if got != want {
			t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
		}
	})

	t.Run("error - ExecuteQuery", func(t *testing.T) {
		mock := mockStore{scenario: Error}
		want := Page{}

		got, err := GetByOID(tOID, mock)

		if err == nil {
			t.Errorf("expected error, got nil")
			t.FailNow()
		}

		if err.Error() != tErrorMsg {
			t.Errorf("expected error %q, got %q", tErrorMsg, err.Error())
			t.FailNow()
		}

		if got != want {
			t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
		}
	})

	t.Run("error - wrong types in columns", func(t *testing.T) {
		// Each column return value with invalid type
		mockes := map[string]mockStore{
			"CreatedOn": {
				scenario:  SuccessGet,
				createdOn: false,
				updatedOn: &tPage.Meta.UpdatedOn,
				dataJSON:  &tDataJSON,
			},
			"UpdatedOn": {
				scenario:  SuccessGet,
				createdOn: &tPage.Meta.CreatedOn,
				updatedOn: false,
				dataJSON:  &tDataJSON,
			},
			"Data": {
				scenario:  SuccessGet,
				createdOn: &tPage.Meta.CreatedOn,
				updatedOn: &tPage.Meta.UpdatedOn,
				dataJSON:  false,
			},
		}

		for columnName, mock := range mockes {
			want := Page{}

			got, err := GetByOID(tOID, mock)

			if err == nil {
				t.Errorf("expected error, got nil")
				t.FailNow()
			}

			wantErr := fmt.Sprintf(typeAssertError, columnName)
			if err.Error() != wantErr {
				t.Errorf("expected error %q, got %q", wantErr, err.Error())
				t.FailNow()
			}

			if got != want {
				t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
			}
		}
	})

	t.Run("error - unmarshaling JSON", func(t *testing.T) {
		mock := mockStore{
			scenario:  SuccessGet,
			createdOn: &tPage.Meta.CreatedOn,
			updatedOn: &tPage.Meta.UpdatedOn,
			dataJSON:  &tDataInvalidJSON,
		}
		want := Page{}

		got, err := GetByOID(tOID, mock)

		if err == nil {
			t.Errorf("expected error, got nil")
			t.FailNow()
		}

		if !strings.HasPrefix(err.Error(), tJSONErrorMsg) {
			t.Errorf("expected error %q, got %q", tJSONErrorMsg, err.Error())
			t.FailNow()
		}

		if got != want {
			t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
		}
	})
}
func TestSave(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := mockStore{
			scenario:  SuccessSave,
			oid:       &tOID,
			createdOn: &tPage.Meta.CreatedOn,
			updatedOn: &tPage.Meta.UpdatedOn,
		}
		want := tPage

		got, err := Save(tPage.Data, mock)

		if err != nil {
			t.Errorf("expected Page, got err: %q", err)
			t.FailNow()
		}

		if got != want {
			t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
		}
	})

	t.Run("error - ExecuteQuery", func(t *testing.T) {
		mock := mockStore{scenario: Error}
		want := Page{}

		got, err := Save(tPage.Data, mock)

		if err == nil {
			t.Errorf("expected error, got nil")
			t.FailNow()
		}

		if err.Error() != tErrorMsg {
			t.Errorf("expected error %q, got %q", tErrorMsg, err.Error())
			t.FailNow()
		}

		if got != want {
			t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
		}
	})

	t.Run("error - wrong types in columns", func(t *testing.T) {
		// Each column return value with invalid type
		mockes := map[string]mockStore{
			"OID": {
				scenario:  SuccessSave,
				oid:       false,
				createdOn: &tPage.Meta.CreatedOn,
				updatedOn: &tPage.Meta.UpdatedOn,
			},
			"CreatedOn": {
				scenario:  SuccessSave,
				oid:       &tOID,
				createdOn: false,
				updatedOn: &tPage.Meta.UpdatedOn,
			},
			"UpdatedOn": {
				scenario:  SuccessSave,
				oid:       &tOID,
				createdOn: &tPage.Meta.CreatedOn,
				updatedOn: false,
			},
		}

		for columnName, mock := range mockes {
			want := Page{}

			got, err := Save(tPage.Data, mock)

			if err == nil {
				t.Errorf("expected error, got nil")
				t.FailNow()
			}

			wantErr := fmt.Sprintf(typeAssertError, columnName)
			if err.Error() != wantErr {
				t.Errorf("expected error %q, got %q", wantErr, err.Error())
				t.FailNow()
			}

			if got != want {
				t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
			}
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := mockStore{
			scenario:  SuccessUpdate,
			oid:       &tOID,
			createdOn: &tPage.Meta.CreatedOn,
			updatedOn: &tPage.Meta.UpdatedOn,
		}
		want := tPage

		got, err := Update(tOID, tPage.Data, mock)

		if err != nil {
			t.Errorf("expected Page, got err: %q", err)
			t.FailNow()
		}

		if got != want {
			t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
		}
	})

	t.Run("error - ExecuteQuery", func(t *testing.T) {
		mock := mockStore{scenario: Error}
		want := Page{}

		got, err := Update(tOID, tPage.Data, mock)

		if err == nil {
			t.Errorf("expected error, got nil")
			t.FailNow()
		}

		if err.Error() != tErrorMsg {
			t.Errorf("expected error %q, got %q", tErrorMsg, err.Error())
			t.FailNow()
		}

		if got != want {
			t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
		}
	})

	t.Run("error - wrong types in columns", func(t *testing.T) {
		// Each column return value with invalid type
		mockes := map[string]mockStore{
			"CreatedOn": {
				scenario:  SuccessUpdate,
				createdOn: false,
				updatedOn: &tPage.Meta.UpdatedOn,
			},
			"UpdatedOn": {
				scenario:  SuccessUpdate,
				createdOn: &tPage.Meta.CreatedOn,
				updatedOn: false,
			},
		}

		for columnName, mock := range mockes {
			want := Page{}

			got, err := Update(tOID, tPage.Data, mock)

			if err == nil {
				t.Errorf("expected error, got nil")
				t.FailNow()
			}

			wantErr := fmt.Sprintf(typeAssertError, columnName)
			if err.Error() != wantErr {
				t.Errorf("expected error %q, got %q", wantErr, err.Error())
				t.FailNow()
			}

			if got != want {
				t.Errorf("\nwant: %+v\ngot: %+v\n", want, got)
			}
		}
	})
}
