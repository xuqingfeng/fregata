// Package wechat is used to send WeChat messages
package wechat

type Config struct {
	Enabled     bool `toml:"enabled"`
	BaseRequest baseRequest
	PassTicket  string
	From        string
	To          string
}

func NewConfig() Config {

	return Config{}
}

func (c Config) Validate() error {

	return nil
}
