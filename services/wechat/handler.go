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
	Text string `json:"text"`
	// TODO 'fregata' group
	To string `json:"to"`
}

func sendMessage(c Config, m message) error {

	return sendWechatMessage(c.BaseRequest, c.PassTicket, m.Text, c.From, m.To)
}
