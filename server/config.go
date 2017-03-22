// package server provide a http server
package server

import (
	"github.com/xuqingfeng/fregata/logging"
	"github.com/xuqingfeng/fregata/services/macos"
	"github.com/xuqingfeng/fregata/services/slack"
	"github.com/xuqingfeng/fregata/services/telegram"
	"github.com/xuqingfeng/fregata/services/wechat"
)

type Config struct {
	Logging  logging.Config  `toml:"logging"`
	Slack    slack.Config    `toml:"slack"`
	Macos    macos.Config    `toml:"macos"`
	Telegram telegram.Config `toml:"telegram"`
	Wechat   wechat.Config   `toml:"wechat"`
}

func NewConfig() *Config {

	c := &Config{}
	c.Logging = logging.NewConfig()
	c.Slack = slack.NewConfig()
	c.Macos = macos.NewConfig()
	c.Telegram = telegram.NewConfig()
	c.Wechat = wechat.NewConfig()

	return c
}

// Validate make sure all service have valid configs
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
	if err := c.Wechat.Validate(); err != nil {
		return err
	}

	return nil
}
