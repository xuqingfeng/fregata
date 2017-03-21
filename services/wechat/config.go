package wechat

type Config struct {
	Enabled     bool `toml:"enabled"`
	BaseRequest baseRequest
	PassTicket  string
}

func NewConfig() Config {

	return Config{}
}

func (c Config) Validate() error {

	return nil
}
