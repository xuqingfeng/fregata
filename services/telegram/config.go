package telegram

import (
	"github.com/pkg/errors"
	"github.com/xuqingfeng/fregata/vars"
)

type Config struct {
	Enabled               bool   `toml:"enabled"`
	URL                   string `toml:"url"`
	Token                 string `toml:"token"`
	ChatId                string `toml:"chat-id"`
	ParseMode             string `toml:"parse-mode"`
	DisableWebPagePreview bool   `toml:"disable-web-page-preview"`
	DisableNotification   bool   `toml:"disable-notification"`
}

func NewConfig() Config {

	return Config{
		URL: vars.TelegramDefaultURL,
	}
}

func (c Config) Validate() error {

	if c.ParseMode != "Markdown" && c.ParseMode != "HTML" {
		return errors.Errorf("parse-mode %s not valid, use 'Markdown' or 'HTML'", c.ParseMode)
	}

	return nil
}
