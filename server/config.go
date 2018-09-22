// Package server provide a http server.
package server

import (
	"github.com/xuqingfeng/fregata/http"
	"github.com/xuqingfeng/fregata/logging"
	"github.com/xuqingfeng/fregata/services/macos"
	"github.com/xuqingfeng/fregata/services/slack"
	"github.com/xuqingfeng/fregata/services/smtp"
	"github.com/xuqingfeng/fregata/services/telegram"
	"github.com/xuqingfeng/fregata/services/twilio"
	"github.com/xuqingfeng/fregata/services/wechat"
)

type Config struct {
	HTTP     http.Config     `toml:"http"`
	Logging  logging.Config  `toml:"logging"`
	Slack    slack.Config    `toml:"slack"`
	Macos    macos.Config    `toml:"macos"`
	Telegram telegram.Config `toml:"telegram"`
	Twilio   twilio.Config   `toml:"twilio"`
	Wechat   wechat.Config   `toml:"wechat"`
	SMTP     smtp.Config     `toml:"smtp"`
}

func NewConfig() *Config {

	c := &Config{}
	c.HTTP = http.NewConfig()
	c.Logging = logging.NewConfig()
	c.Slack = slack.NewConfig()
	c.Macos = macos.NewConfig()
	c.Telegram = telegram.NewConfig()
	c.Twilio = twilio.NewConfig()
	c.Wechat = wechat.NewConfig()
	c.SMTP = smtp.NewConfig()

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
	if err := c.Twilio.Validate(); err != nil {
		return err
	}
	if err := c.Wechat.Validate(); err != nil {
		return err
	}
	if err := c.SMTP.Validate(); err != nil {
		return err
	}

	return nil
}
