package page

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockStore struct {
	scenario int
	oid      interface{}
}

func (q mockStore) ExecuteQuery(queryCode string, params []interface{}, dest ...interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("unexpected end for mock")
}

func TestHandleGET(t *testing.T) {
	t.Run("error - no OID", func(t *testing.T) {
		rq, _ := http.NewRequest(http.MethodGet, "/api/v1/page", nil)
		rs := httptest.NewRecorder()

		HandleGET(rs, rq)

		got := rs.Body.String()
		want := invalidOID

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("error - invalid OID", func(t *testing.T) {
		rq, _ := http.NewRequest(http.MethodGet, "/api/v1/page?oid=123", nil)
		rs := httptest.NewRecorder()

		HandleGET(rs, rq)

		got := rs.Body.String()
		want := invalidOID

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	// t.Run("error - GetByOID", func(t *testing.T) {
	// 	rq, _ := http.NewRequest(http.MethodGet, "/api/v1/page?oid=bb7a2a6a-66e0-4680-ba0f-bec2d3764e68", nil)
	// 	rs := httptest.NewRecorder()

	// 	HandleGET(rs, rq)

	// 	got := rs.Body.String()
	// 	want := invalidOID

	// 	if got != want {
	// 		t.Errorf("want %q, got %q", want, got)
	// 	}
	// })
}
