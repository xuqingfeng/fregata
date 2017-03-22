package services

import (
	"encoding/json"
	"net/http"
)

// Msg is the general message returned to client
type Msg struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SendMessage send json data to client
func SendMessage(msg Msg, status int, w http.ResponseWriter) {

	msgInJSON, _ := json.Marshal(msg)
	w.WriteHeader(status)
	w.Write(msgInJSON)
}
