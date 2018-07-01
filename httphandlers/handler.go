package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../storage"
	"../structs/httprequest"
	"./httputil"
)

func GetPolls(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	httputil.HandleSuccess(&w, storage.ListPolls())
}

func GetPollOptions(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	var req httprequest.GetPollOptionsRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		httputil.HandleError(&w, 400, "Bad Request", "Error marshalling request JSON", err)
	} else {
		if poll, ok := storage.GetPoll(req.PollID); ok {
			httputil.HandleSuccess(&w, poll)
		} else {
			httputil.HandleError(&w, 400, "Bad Request", "Entity does not exist", err)
		}
	}
}

func SelectOption(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	decoder := json.NewDecoder(r.Body)
	var req httprequest.SelectOptionRequest
	err := decoder.Decode(&req)
	if err != nil {
		httputil.HandleError(&w, 400, "Bad Request", "Error marshalling request JSON", err)
	} else {
		if ok := storage.SelectPollOption(req.PollID, req.OptionID); ok {
			httputil.HandleSuccess(&w, ok)
		} else {
			httputil.HandleError(&w, 400, "Bad Request", "Entity does not exist", err)
		}
	}
}

func AddPoll(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	decoder := json.NewDecoder(r.Body)
	var req httprequest.AddPollRequest
	err := decoder.Decode(&req)
	if err != nil {
		httputil.HandleError(&w, 400, "Bad Request", "Error marshalling request JSON", err)
	} else {
		if pollID := storage.AddPoll(req.Question, req.OptionList); pollID > 0 {
			httputil.HandleSuccess(&w, pollID)
		} else {
			httputil.HandleError(&w, 400, "Bad Request", "Entity does not exist", err)
		}
	}
}

func AddOption(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	decoder := json.NewDecoder(r.Body)
	var req httprequest.AddOptionRequest
	err := decoder.Decode(&req)
	if err != nil {
		httputil.HandleError(&w, 400, "Bad Request", "Error marshalling request JSON", err)
	} else {
		if optionID, ok := storage.AddOption(req.Option, req.PollID); ok {
			httputil.HandleSuccess(&w, optionID)
		} else {
			httputil.HandleError(&w, 400, "Bad Request", "Entity does not exist", err)
		}
	}
}

func RemovePoll(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	decoder := json.NewDecoder(r.Body)
	var req httprequest.RemovePollRequest
	err := decoder.Decode(&req)
	if err != nil {
		httputil.HandleError(&w, 400, "Bad Request", "Error marshalling request JSON", err)
	} else {
		if ok := storage.RemovePoll(req.PollID); ok {
			httputil.HandleSuccess(&w, ok)
		} else {
			httputil.HandleError(&w, 400, "Bad Request", "Entity does not exist", err)
		}
	}
}
func RemoveOption(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming Request:", r)
	decoder := json.NewDecoder(r.Body)
	var req httprequest.RemoveOptionRequest
	err := decoder.Decode(&req)
	if err != nil {
		httputil.HandleError(&w, 400, "Bad Request", "Error marshalling request JSON", err)
	} else {
		if ok := storage.RemoveOption(req.PollID, req.OptionID); ok {
			httputil.HandleSuccess(&w, ok)
		} else {
			httputil.HandleError(&w, 400, "Bad Request", "Entity does not exist", err)
		}
	}
}
