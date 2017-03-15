package slack

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
	Text        string       `json:"text"`
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []attachment `json:"attachments"`
}

type attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	Pretext    string  `json:"pretext"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Fields     []field `json:"fields"`
	ImageUrl   string  `json:"image_url"`
	ThumbUrl   string  `json:"thumb_url"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	Ts         int     `json:"ts"`
}

type field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func sendMessage(c Config, m message) error {

	if m.Channel == "" {
		m.Channel = c.Channel
	}

	dataInJSON, err := json.Marshal(m)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(dataInJSON)

	resp, err := http.Post(c.URL, "application/json", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		type errResponse struct {
			Error string `json:"error"`
		}
		e := new(errResponse)
		err = json.NewDecoder(resp.Body).Decode(e)
		if err != nil {
			return err
		}
		return errors.New(e.Error)
	}

	return nil
}
