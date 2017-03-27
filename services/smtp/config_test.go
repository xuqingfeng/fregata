package smtp

import "testing"

func TestValidate(t *testing.T) {

	tests := []struct {
		c         Config
		shouldErr bool
	}{
		{Config{Host: "http://192.168.0.%31"}, true},
		{Config{Host: "localhost", Port: 25}, false},
		{Config{Host: "https://example.com"}, false},
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
