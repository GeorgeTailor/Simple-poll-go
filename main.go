package main

import (
	"log"
	"net/http"

	"./httphandlers"
	"./storage"
	"./structs"
	"github.com/gorilla/mux"
)

const PORT = 8080

var messageId = 0

func createOption(option string) structs.Option {
	messageId++
	return structs.Option{
		ID:     messageId,
		Option: option,
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/poll_options", httphandlers.GetPollOptions).Methods("GET")
	router.HandleFunc("/select_option/{id}", httphandlers.SelectOption).Methods("POST")
	router.HandleFunc("/poll_results", httphandlers.GetPollResults).Methods("GET")

	storage.Add(createOption("Testing"))
	storage.Add(createOption("Testing Again"))
	storage.Add(createOption("Testing A Third Time"))

	log.Fatal(http.ListenAndServe(":8080", router))

	// log.Println("Attempting to start HTTP Server.")

	// http.HandleFunc("/", httphandlers.HandleRequest)

	// var err = http.ListenAndServe(":"+strconv.Itoa(PORT), nil)

	// if err != nil {
	// 	log.Panicln("Server failed starting. Error: %s", err)
	// }
}
