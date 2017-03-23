package main

import "testing"

func TestParseConfig(t *testing.T) {

	tests := []struct {
		path      string
		shouldErr bool
	}{
		{"", true},
		{"testdata/fregata.good.conf", false},
		{"testdata/fregata.bad.conf", true},
	}

	for i, test := range tests {
		_, err := parseConfig(test.path)
		if err != nil && !test.shouldErr {
			t.Errorf("E! test %d shouldn't return error got %s", i, err.Error())
		}
		if err == nil && test.shouldErr {
			t.Errorf("E! test %d should return error", i)
		}
	}

}
