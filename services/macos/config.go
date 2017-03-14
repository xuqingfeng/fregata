package macos

import (
	"github.com/pkg/errors"
	"runtime"
)

type Config struct {
	Enabled bool `toml:"enabled"`
}

func NewConfig() Config {

	return Config{}
}

func (c Config) Validate() error {

	if runtime.GOOS != "darwin" {
		return errors.New("os is not macOS")
	}

	return nil
}
