package twilio

import "testing"

func TestValidate(t *testing.T) {

	tests := []struct {
		c         Config
		shouldErr bool
	}{
		{Config{Enabled: true, AccountSid: "TRSLoL4oAsbwFFuIZz6xfwo0aHqpss9H5Z", AuthToken: "pQTUqSCr2vaokZeSdjp2rOuPxvi0czms", From: "+123456"}, false},
		{Config{Enabled: true, AccountSid: "YG3X2H01IoMCYys5Kaf56yXzMKt4TAv9"}, true},
		{Config{Enabled: true, AccountSid: "RDDk51lPk19xfxvUzYA3EBzDV5c16faWqC"}, true},
		{Config{From: "123456"}, true},
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
