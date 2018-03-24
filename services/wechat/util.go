// Coding according to this document: https://github.com/biezhi/wechat-robot/blob/master/doc/protocol.md
// http://www.sunrui123.com/archives/15
package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/xuqingfeng/fregata/vars"
)

// Login is used to ask user to scan QR code and obtain keys for later use.
func (s *Service) Login() (baseRequest, string, error) {

	var b baseRequest
	uuid, err := s.getUUID()
	if err != nil {
		return b, "", err
	}
	err = s.getQR(uuid)
	if err != nil {
		return b, "", err
	}
	redirectUrl, err := s.waitForLogin(uuid, 1)
	if err != nil {
		return b, "", err
	}
	keys, err := s.getKeys(redirectUrl)
	if err != nil {
		return b, "", err
	}
	b = baseRequest{
		Uin:      keys.Wxuin,
		Sid:      keys.Wxsid,
		Skey:     keys.Skey,
		DeviceID: s.getDeviceID(),
	}

	return b, keys.PassTicket, nil
}

func (s *Service) getUUID() (string, error) {

	params := url.Values{}
	params.Set("appid", vars.WechatAppID)
	params.Set("fun", vars.WechatFun)
	params.Set("lang", vars.WechatLang)
	params.Set("_", strconv.FormatInt(time.Now().Unix(), 10))

	resp, err := http.Post(vars.WechatUUIDUrl, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`window.QRLogin.code = (\d+); window.QRLogin.uuid = "(\S+?)";`)
	matches := re.FindStringSubmatch(string(body))
	if code := matches[1]; code != "200" {
		return "", errors.Errorf("get uuid status code: %s", code)
	} else if uuid := matches[2]; uuid == "" {
		return "", errors.New("uuid empty")
	}

	return matches[2], nil
}

func (s *Service) getQR(uuid string) error {

	// TODO: 17/3/20 output QR code in terminal
	qrcodeUrl := fmt.Sprintf("%s/%s", vars.WechatQRUrl, uuid)
	// always print this message
	fmt.Printf("[wechat] scan the QR code(%s) and login.\n", qrcodeUrl)

	return nil
}

const (
	fg = "\033[48;5;2m  \033[0m"
	bg = "\033[48;5;7m  \033[0m"
)

// waitForLogin wait user scan QR code and login
func (s *Service) waitForLogin(uuid string, tip int) (string, error) {

	var redirectUrl string
	retry := 5
Wait:
	for {
		if retry <= 0 {
			break
		}
		resp, err := http.Get(fmt.Sprintf("%s?tip=%d&uuid=%s&_=%d", vars.WechatLoginUrl, tip, uuid, time.Now().Unix()))
		if err != nil {
			return "", err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		re := regexp.MustCompile(`window.code=(\d+);`)
		matches := re.FindStringSubmatch(string(body))

		code := matches[1]

		switch code {
		case "200":
			s.logger.Printf("I! login success %s", code)
			re = regexp.MustCompile(`window.redirect_uri="(\S+?)";`)
			matches = re.FindStringSubmatch(string(body))
			redirectUrl = matches[1] + "&fun=" + vars.WechatFun
			resp.Body.Close()
			break Wait
		case "201":
			s.logger.Printf("I! scan success %s", code)
		case "408":
			s.logger.Printf("I! timeout %s, please restart", code)
		}
		resp.Body.Close()
		retry--
		time.Sleep(5 * time.Second)
	}

	return redirectUrl, nil
}

type redirectUrlResp struct {
	XMLName     xml.Name `xml:"error"`
	Ret         int      `xml:"ret"`
	Message     string   `xml:"message"`
	Skey        string   `xml:"skey"`
	Wxsid       string   `xml:"wxsid"`
	Wxuin       int64    `xml:"wxuin"`
	PassTicket  string   `xml:"pass_ticket"`
	Isgrayscale int      `xml:"isgrayscale"`
}

func (s *Service) getKeys(redirectUrl string) (redirectUrlResp, error) {

	var ret redirectUrlResp

	resp, err := http.Get(redirectUrl)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()

	err = xml.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// getDeviceID generate fake device id
func (s *Service) getDeviceID() string {

	base := []byte("0123456789")
	var deviceID []byte
	rand.Seed(15)
	for i := 0; i < 15; i++ {
		deviceID = append(deviceID, base[rand.Intn(len(base))])
	}

	return "e" + string(deviceID)
}

type baseRequest struct {
	Uin      int64  `json:"Uin"`
	Sid      string `json:"Sid"`
	Skey     string `json:"Skey"`
	DeviceID string `json:"DeviceID"`
}
type baseResponse struct {
	Ret    int    `json:"Ret"`
	ErrMsg string `json:"ErrMsg"`
}
type syncKey struct {
	Count int    `json:"Count"`
	List  []list `json:"List"`
}
type list struct {
	Key   int `json:"Key"`
	Value int `json:"Value"`
}
type user struct {
	Uin      int64  `json:"Uin"`
	UserName string `json:"UserName"`
	NickName string `json:"NickName"`
}

func (s *Service) wxInit(b baseRequest, pass_ticket string) (string, error) {
	// probably don't need this

	type params struct {
		BaseRequest baseRequest `json:"BaseRequest"`
	}
	p := params{
		b,
	}
	dataInJSON, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(dataInJSON)

	resp, err := http.Post(vars.WechatInitUrl+"?pass_ticket="+pass_ticket, "application/json", reader)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", errors.New(string(body))
	}

	type initResponse struct {
		BaseResponse baseResponse `json:"BaseResponse"`
		SyncKey      syncKey      `json:"SyncKey"`
		User         user         `json:"User"`
	}

	var r initResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	if r.BaseResponse.Ret != 0 {
		return "", errors.Errorf("Ret: %d, ErrMsg: %s", r.BaseResponse.Ret, r.BaseResponse.ErrMsg)
	}

	return r.User.UserName, nil
}

// getContact will return "fregata" group username if group exits and are saved to contact list
// else return empty string
func (s *Service) getContact(b baseRequest, pass_ticket string) (string, error) {

	type params struct {
		BaseRequest baseRequest `json:"BaseRequest"`
	}
	p := params{
		b,
	}
	dataInJSON, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(vars.WechatGetContactUrl+"?pass_ticket="+pass_ticket, "application/json", bytes.NewReader(dataInJSON))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", errors.New(string(body))
	}

	type getContactResponse struct {
		BaseResponse baseResponse `json:"BaseResponse"`
		MemberCount  int          `json:"MemberCount"`
		MemberList   []user       `json:"MemberList"`
	}
	var r getContactResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}
	if r.BaseResponse.Ret != 0 {
		return "", errors.Errorf("Ret: %d, ErrMsg: %s", r.BaseResponse.Ret, r.BaseResponse.ErrMsg)
	}

	for _, u := range r.MemberList {
		if name := strings.ToLower(u.NickName); name == "fregata" {
			return u.UserName, nil
		}
	}

	return "", nil
}

// batchGetContact try to get 'fregata' group in un-saved contact list
// TODO: 2017/3/31 https://github.com/HalfdogStudio/wechat-user-bot/issues/17
func (s *Service) batchGetContact(b baseRequest, pass_ticket string) (string, error) {

	type params struct {
	}
	return "", nil
}

func (s *Service) notify(b baseRequest, pass_ticket, from, to string) error {

	type params struct {
		BaseRequest  baseRequest `json:"BaseRequest"`
		Code         int         `json:"Code"`
		FromUserName string      `json:"FromUserName"`
		ToUserName   string      `json:"ToUserName"`
		ClientMsgId  int         `json:"ClientMsgId"`
	}

	p := params{
		BaseRequest:  b,
		Code:         3,
		FromUserName: from,
		ToUserName:   to,
		ClientMsgId:  int(time.Now().Unix() * 1e4),
	}
	dataInJSON, err := json.Marshal(p)
	if err != nil {
		return err

	}
	resp, err := http.Post(vars.WechatNotifyUrl+"?pass_ticket="+pass_ticket, "application/json", bytes.NewReader(dataInJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	type notifyResponse struct {
		BaseResponse baseResponse `json:"BaseResponse"`
	}

	var r notifyResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return err
	}
	if r.BaseResponse.Ret != 0 {
		return errors.Errorf("Ret: %d, ErrMsg: %s", r.BaseResponse.Ret, r.BaseResponse.ErrMsg)
	}

	return nil
}

func sendWechatMessage(b baseRequest, pass_ticket, content, from, to string) error {

	type msg struct {
		Type         int    `json:"Type"`
		Content      string `json:"Content"`
		FromUserName string `json:"FromUserName"`
		ToUserName   string `json:"ToUserName"`
		LocalID      int    `json:"LocalID"`
		ClientMsgId  int    `json:"ClientMsgId"`
	}
	type params struct {
		BaseRequest baseRequest `json:"BaseRequest"`
		Msg         msg         `json:"Msg"`
	}
	p := params{
		BaseRequest: b,
		Msg: msg{
			Type:         1,
			Content:      content,
			FromUserName: from,
			ToUserName:   to,
			LocalID:      int(time.Now().Unix() * 1e4),
			ClientMsgId:  int(time.Now().Unix() * 1e4),
		},
	}
	dataInJSON, err := json.Marshal(p)
	if err != nil {
		return err
	}
	resp, err := http.Post(vars.WechatSendMessageUrl+"?pass_ticket="+pass_ticket, "application/json", bytes.NewReader(dataInJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	type sendMessageResponse struct {
		BaseResponse baseResponse `json:"BaseResponse"`
	}

	var r sendMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return err
	}
	if r.BaseResponse.Ret != 0 {
		return errors.Errorf("Ret: %d, ErrMsg: %s", r.BaseResponse.Ret, r.BaseResponse.ErrMsg)
	}

	return nil
}
