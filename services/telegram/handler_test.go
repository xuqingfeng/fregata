package telegram

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
		{Config{ParseMode: "markdown"}, message{}, true},
		{Config{}, message{ParseMode: "html"}, true},
		{Config{}, message{ParseMode: "html"}, true},
		{Config{URL: "https://example.com"}, message{ParseMode: "HTML"}, true}, // no token
		{Config{URL: testServer.URL + "/", Token: "test-token", ParseMode: "Markdown"}, message{Message: "test"}, false},
	}

	//t.Logf("I! testServer url %s", testServer.URL)
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
