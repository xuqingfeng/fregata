// Package vars keeps all variables.
package vars

const (
	/*
	 * App
	 */
	Name = "fregata"

	/*
	 * API
	 */
	APIBasePath = "/" + Name + "/v1"

	/*
	 * Service
	 */
	// slack
	SlackDefaultUsername = Name

	// telegram
	TelegramDefaultURL = "https://api.telegram.org/bot"

	// wechat
	WechatAppID          = "wx782c26e4c19acffb"
	WechatFun            = "new"
	WechatLang           = "zh_CN"
	WechatUUIDUrl        = "https://login.weixin.qq.com/jslogin"
	WechatQRUrl          = "https://login.weixin.qq.com/qrcode"
	WechatLoginUrl       = "https://login.weixin.qq.com/cgi-bin/mmwebwx-bin/login"
	WechatInitUrl        = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxinit"
	WechatGetContactUrl  = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxgetcontact"
	WechatNotifyUrl      = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxstatusnotify"
	WechatSendMessageUrl = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsendmsg"
)
