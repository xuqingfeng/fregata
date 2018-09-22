package twilio

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
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
	From string `json:"from"`
	Text string `json:"text"`
	To   string `json:"to"`
}

func sendMessage(c Config, m message) error {

	if m.From == "" {
		m.From = c.From
	}

	formValues := url.Values{}
	formValues.Set("From", m.From)
	formValues.Set("To", m.To)
	formValues.Set("Body", m.Text)
	req, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/"+c.AccountSid+"/Messages.json", strings.NewReader(formValues.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.AccountSid, c.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		type errorResponse struct {
			Sid          string `json:"sid"`
			AccountSid   string `json:"account_sid"`
			From         string `json:"from"`
			To           string `json:"to"`
			Body         string `json:"body"`
			Status       string `json:"status"`
			NumSegments  string `json:"num_segments"`
			ErrorCode    int    `json:"error_code"`
			ErrorMessage string `json:"error_message"`
		}
		e := new(errorResponse)
		err = json.NewDecoder(resp.Body).Decode(e)
		if err != nil {
			return err

		}

		return errors.Errorf("%v", e)
	}

	return nil
}
