package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../storage"
	"../structs"
	"./httputil"
)

func GetPollOptions(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	httputil.HandleSuccess(&w, storage.GetPollOptions())
}

func SelectOption(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	decoder := json.NewDecoder(r.Body)
	var option structs.Option
	err := decoder.Decode(&option)
	if err != nil {
		httputil.HandleError(&w, 400, "Bad Request", "Error marshalling request JSON", err)
		return
	} else {
		if opt, ok := storage.SelectPollOption(option.ID); ok {
			httputil.HandleSuccess(&w, opt)
		} else {
			httputil.HandleError(&w, 400, "Bad Request", "Entity does not exist", err)
		}
	}
}

func GetPollResults(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)

}
