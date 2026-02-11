package slack

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMessage(t *testing.T) {

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	testServer := httptest.NewServer(http.HandlerFunc(testHandler))
	defer testServer.Close()

	tests := []struct {
		c         Config
		m         message
		shouldErr bool
	}{
		{Config{URL: "https://test.test"}, message{Message: "test"}, true}, // unreachable host
		{Config{URL: testServer.URL, Channel: "test"}, message{Message: "test"}, false},
	}

	for i, test := range tests {
		err := sendMessage(test.c, test.m)
		if err != nil && !test.shouldErr {
			t.Errorf("E! test %d shouldn't return error got %s", i, err.Error())
		}
		if err == nil && test.shouldErr {
			t.Errorf("E! test %d should return error", i)
		}
	}

}
