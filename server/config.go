package server

import (
	"github.com/xuqingfeng/fregata/logging"
	"github.com/xuqingfeng/fregata/services/macos"
	"github.com/xuqingfeng/fregata/services/slack"
	"github.com/xuqingfeng/fregata/services/telegram"
)

type Config struct {
	Logging  logging.Config  `toml:"logging"`
	Slack    slack.Config    `toml:"slack"`
	Macos    macos.Config    `toml:"macos"`
	Telegram telegram.Config `toml:"telegram"`
}

func NewConfig() *Config {

	c := &Config{}
	c.Logging = logging.NewConfig()
	c.Slack = slack.NewConfig()
	c.Macos = macos.NewConfig()
	c.Telegram = telegram.NewConfig()

	return c
}

func (c *Config) Validate() error {

	if err := c.Slack.Validate(); err != nil {
		return err
	}
	if err := c.Macos.Validate(); err != nil {
		return err
	}
	if err := c.Telegram.Validate(); err != nil {
		return err
	}

	return nil
}
