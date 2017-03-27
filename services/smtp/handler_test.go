package smtp

import "testing"

func TestSendMessage(t *testing.T) {

	tests := []struct {
		c         Config
		m         message
		shouldErr bool
	}{
		{Config{}, message{From: ""}, true},
		{Config{}, message{From: "fregata@localhost", To: []string{}}, true},
		{Config{From: ""}, message{From: "fregata@localhost"}, true},
		{Config{Host: "", Port: 25}, message{From: "fregata@localhost", To: []string{"fregata@example.com"}}, true},
		//{Config{Host: "localhost", Port: 25}, message{From: "fregata@localhost", To: []string{"fregata@example.com"}}, false},
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
