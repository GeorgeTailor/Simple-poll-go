package httphandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log"

	"../storage"
	"../structs"
	"./httputil"
)

func Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	byteData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httputil.HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
		return
	}

	var option structs.Option

	err = json.Unmarshal(byteData, &option)

	if err != nil {
		httputil.HandleError(&w, 500, "Internal Server Error", "Error unmarhsalling JSON", err)
		return
	}

	if option.Option == "" {
		httputil.HandleError(&w, 400, "Bad Request", "Unmarshalled JSON didn't have required fields", nil)
		return
	}

	id := storage.Add(option)

	log.Println("Added option:", option)

	httputil.HandleSuccess(&w, structs.ID{ID: id})
}
