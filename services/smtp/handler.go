package smtp

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xuqingfeng/fregata/services"
	"gopkg.in/gomail.v2"
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
	From    string   `json:"from"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func sendMessage(c Config, m message) error {

	if m.From == "" {
		m.From = c.From
	}
	if len(m.To) == 0 {
		m.To = c.To
	}

	if len(m.From) == 0 {
		return errors.New("from is empty")
	}
	if len(m.To) == 0 {
		return errors.New("to is empty")
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", m.From)
	mail.SetHeader("To", m.To...)
	if len(m.Cc) != 0 {
		mail.SetHeader("Cc", m.Cc...)
	}
	mail.SetHeader("Subject", m.Subject)
	mail.SetBody("text/html", m.Body)

	d := gomail.NewDialer(c.Host, c.Port, c.Username, c.Password)
	if err := d.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
