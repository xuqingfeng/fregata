// Package twilio is used to send SMS through twilio platform
package twilio

import (
	"strings"

	"github.com/pkg/errors"
)

type Config struct {
	Enabled    bool   `toml:"enabled"`
	AccountSid string `toml:"account-sid"`
	AuthToken  string `toml:"auth-token"`
	From       string `toml:"from"`
}

func NewConfig() Config {

	return Config{}
}

func (c Config) Validate() error {

	if c.Enabled {
		if len(c.AccountSid) != 34 {
			return errors.New("account-sid must have 34 characters")
		}

		if len(c.AuthToken) == 0 {
			return errors.New("auth-token can't be empty")
		}
	}

	/*
	  from can be defined later in API call
	  but if it's defined in config file, it must have the right format
	*/
	if len(c.From) != 0 && !strings.HasPrefix(c.From, "+") {
		// https://www.twilio.com/docs/glossary/what-e164
		return errors.Errorf("phone number(%s) is not valid: [+][country code][subscriber number including area code]", c.From)
	}

	return nil
}
