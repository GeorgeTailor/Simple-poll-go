package httphandlers

import "net/http"
import "log"
import "./httputil"
import "../storage"

func GetPollOptions(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	httputil.HandleSuccess(&w, storage.GetPollOptions())
}

func SelectOption(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)

}

func GetPollResults(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)

}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	switch r.Method {
	case http.MethodGet:
		List(w, r)
		break
	case http.MethodPost:
		Add(w, r)
		break
	case http.MethodDelete:
		Remove(w, r)
		break
	default:
		httputil.HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}
