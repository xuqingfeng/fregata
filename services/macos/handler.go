package macos

import (
	"encoding/json"
	"net/http"

	"github.com/everdev/mack"
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
	Text     string `json:"text"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Sound    string `json:"sound"`
}

func sendMessage(c Config, m message) error {

	return mack.Notify(m.Text, m.Title, m.Subtitle, m.Sound)
}
