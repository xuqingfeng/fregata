package server

import "net/http"

// ServiceHandler return pong to ping
func ServiceHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("pong"))
}
