package telegram

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
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
	Text      string `json:"text"`
	ChatId    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}

func sendMessage(c Config, m message) error {

	if m.ChatId == "" {
		m.ChatId = c.ChatId
	}

	if m.ParseMode == "" {
		m.ParseMode = c.ParseMode
	}
	if m.ParseMode != "Markdown" && m.ParseMode != "HTML" {
		return errors.Errorf("parse-mode %s not valid, use 'Markdown' or 'HTML'", c.ParseMode)
	}

	dataInJSON, err := json.Marshal(m)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(dataInJSON)

	sendMessageUrl := c.URL + c.Token + "/sendMessage"
	resp, err := http.Post(sendMessageUrl, "application/json", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		type errResponse struct {
			Description string `json:"description"`
			ErrorCode   int    `json:"error_code"`
			Ok          bool   `json:"ok"`
		}
		e := new(errResponse)
		err = json.NewDecoder(resp.Body).Decode(e)
		if err != nil {
			return err
		}
		return errors.Errorf("%v", e)
	}

	return nil
}
