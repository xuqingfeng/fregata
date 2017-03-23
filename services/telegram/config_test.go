package telegram

import "testing"

func TestValidate(t *testing.T) {

	tests := []struct {
		c         Config
		shouldErr bool
	}{
		{Config{Enabled: true, URL: "", ParseMode: "Markdown"}, true},
		{Config{URL: "http://192.168.0.%31", ParseMode: "HTML"}, true},
		{Config{URL: "https://example.com", ParseMode: "markdown"}, true},
		{Config{URL: "https://example.com", ParseMode: "html"}, true},
		{Config{URL: "https://example.com", ParseMode: "Markdown"}, false},
		{Config{URL: "https://example.com", ParseMode: "HTML"}, false},
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
