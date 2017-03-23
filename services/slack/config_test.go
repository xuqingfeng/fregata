package slack

import (
	"testing"
)

// When url.Parse will return error? - https://golang.org/src/net/url/url_test.go
func TestValidate(t *testing.T) {

	tests := []struct {
		c         Config
		shouldErr bool
	}{
		{Config{Enabled: true, URL: "http://192.168.0.%31"}, true},
		{Config{Enabled: true, URL: "/"}, false},
		{Config{Enabled: false, URL: "test.test"}, false},
		{Config{Enabled: true, URL: "example.com"}, false},
	}

	for i, test := range tests {
		err := test.c.Validate()
		if err != nil && !test.shouldErr {
			t.Errorf("E! test %d shouldn't return error got %s", i, err.Error())
		}
		if err == nil && test.shouldErr {
			t.Errorf("E! test %d should return error", i)
		}
	}
}
