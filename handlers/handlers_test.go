package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	// a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestAssetsHandler(t *testing.T) {

	req, _ := http.NewRequest("GET", "/coins", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
