package server

import (
	"net/http"

	"github.com/xuqingfeng/fregata/services"
)

// ServiceHandler return pong to ping
func ServiceHandler(w http.ResponseWriter, r *http.Request) {

	msg := services.Msg{
		Success: true,
		Message: "pong",
	}
	services.SendMessage(msg, http.StatusOK, w)
}
