package main

import (
	"log"
	"net/http"

	"./httphandlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/polls", httphandlers.GetPolls).Methods("GET")
	router.HandleFunc("/poll_options", httphandlers.GetPollOptions).Methods("POST")
	router.HandleFunc("/select_option", httphandlers.SelectOption).Methods("POST")
	router.HandleFunc("/add_poll", httphandlers.AddPoll).Methods("POST")
	router.HandleFunc("/add_option", httphandlers.AddOption).Methods("POST")
	router.HandleFunc("/remove_poll", httphandlers.RemovePoll).Methods("DELETE")
	router.HandleFunc("/remove_option", httphandlers.RemoveOption).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
