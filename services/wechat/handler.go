package wechat

import (
	"encoding/json"
	"net/http"

	"github.com/xuqingfeng/fregata/services"
)

func ServiceHandler(c Config) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var m message
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			msg := services.Msg{
				Success: false,
				Message: err.Error(),
			}
			services.SendMessage(msg, http.StatusBadRequest, w)
			return
		}
		err = sendMessage(c, m)
		if err != nil {
			msg := services.Msg{
				Success: false,
				Message: err.Error(),
			}
			services.SendMessage(msg, http.StatusBadRequest, w)
			return
		}
		msg := services.Msg{
			Success: true,
		}
		services.SendMessage(msg, http.StatusOK, w)
	}
}

type message struct {
	Message string `json:"message"`
	To      string `json:"to"`
}

func sendMessage(c Config, m message) error {

	if m.To == "" {
		m.To = c.To
	}
	return sendWechatMessage(c.BaseRequest, c.PassTicket, m.Message, c.From, m.To)
}
