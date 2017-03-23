// Package vars keeps all variables.
package vars

const (
	/*
	   App
	*/
	Name       = "fregata"
	DaemonName = "fregatad"
	Version    = "0.0.1"

	/*
	   API
	*/
	APIBasePath = "/" + Name + "/v1"

	/*
	   Service
	*/
	// slack
	SlackDefaultUsername = Name

	// telegram
	TelegramDefaultURL = "https://api.telegram.org/bot"

	// wechat
	//WechatAppID   = "wxeb7ec651dd0aefa9"
	WechatAppID          = "wx782c26e4c19acffb"
	WechatFun            = "new"
	WechatLang           = "zh_CN"
	WechatUUIDUrl        = "https://login.weixin.qq.com/jslogin"
	WechatQRUrl          = "https://login.weixin.qq.com/qrcode"
	WechatLoginUrl       = "https://login.weixin.qq.com/cgi-bin/mmwebwx-bin/login"
	WechatInitUrl        = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxinit"
	WechatSendMessageUrl = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsendmsg"
)
