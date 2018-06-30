package httphandlers

import (
	"net/http"

	"../storage"
	"./httputil"
)

func List(w http.ResponseWriter, r *http.Request) {
	httputil.HandleSuccess(&w, storage.Get())
}
