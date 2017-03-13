package slack

import "net/http"

func ServiceHandler(w http.ResponseWriter, r *http.Request) {

	// send slack message
	w.Write([]byte("slack"))
}

func sendSlack() {

}
