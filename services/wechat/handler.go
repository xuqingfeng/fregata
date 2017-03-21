package wechat

import (
	"encoding/json"
	"net/http"
)

func ServiceHandler(c Config) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var m message
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = sendMessage(c, m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

type message struct {
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

func sendMessage(c Config, m message) error {

	return sendWechatMessage(c.BaseRequest, c.PassTicket, m.Text, m.From, m.To)
}
