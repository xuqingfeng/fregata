// Package http is only used for configure custom bind address
package http

type Config struct {
	BindAddress string `toml:"bind-address"`
}

func NewConfig() Config {

	return Config{}
}
