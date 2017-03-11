package slack

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/xuqingfeng/fregata/vars"
)

type Config struct {
	Enabled  bool   `toml:"enabled"`
	URL      string `toml:"url"`
	Channel  string `toml:"channel"`
	Username string `toml:"username"`
}

func NewConfig() Config {

	return Config{
		Username: vars.SlackDefaultUsername,
	}
}

func (c Config) Validate() error {

	if c.Enabled && c.URL == "" {
		return errors.New("must specify url")
	}
	if _, err := url.Parse(c.URL); err != nil {
		return errors.Wrapf(err, "invalid url %q", c.URL)
	}

	return nil
}
