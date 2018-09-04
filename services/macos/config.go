// Package macos is used to send notifications to macOS
package macos

import (
	"runtime"

	"github.com/pkg/errors"
)

type Config struct {
	Enabled bool `toml:"enabled"`
}

func NewConfig() Config {

	return Config{}
}

// Validate check if OS is macOS
func (c Config) Validate() error {

	if c.Enabled && runtime.GOOS != "darwin" {
		return errors.New("os is not macOS")
	}

	return nil
}
