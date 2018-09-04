// Package smtp is used to send emails through SMTP
package smtp

import "net/url"

type Config struct {
	Enabled  bool     `toml:"enabled"`
	Host     string   `toml:"host"`
	Port     int      `toml:"port"`
	Username string   `toml:"username"`
	Password string   `toml:"password"`
	From     string   `toml:"from"`
	To       []string `toml:"to"`
}

func NewConfig() Config {

	return Config{}
}

func (c Config) Validate() error {

	if _, err := url.Parse(c.Host); err != nil {
		return err
	}
	return nil
}
