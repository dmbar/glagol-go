package page

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGET(t *testing.T) {
	t.Run("error - no OID", func(t *testing.T) {
		rq, _ := http.NewRequest(http.MethodGet, "/api/v1/page", nil)
		rs := httptest.NewRecorder()

		HandleGET(rs, rq)

		got := rs.Body.String()
		want := "OID was not provided or it is not a valid UUID value"

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})
}
