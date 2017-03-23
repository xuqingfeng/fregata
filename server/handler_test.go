package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceHandler(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(ServiceHandler))
	defer testServer.Close()

	resp, err := http.Get(testServer.URL + "/ping")
	if err != nil {
		t.Errorf("E! get '/ping' fail with %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("E! response not ok, status code: %d ", resp.StatusCode)
	}

}
