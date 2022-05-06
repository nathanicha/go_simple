package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type event struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          1,
		Title:       "Monday",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Image:       "",
	},
	{
		ID:          2,
		Title:       "Tuesday",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Image:       "",
	},
	{
		ID:          3,
		Title:       "Wednesday",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Image:       "",
	},
	{
		ID:          4,
		Title:       "Thursday",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Image:       "",
	},
	{
		ID:          5,
		Title:       "Friday",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Image:       "",
	},
	{
		ID:          6,
		Title:       "Saturday",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Image:       "",
	},
	{
		ID:          7,
		Title:       "Sunday",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
		Image:       "",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home2!")
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if string(rune(singleEvent.ID)) == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func getToday(w http.ResponseWriter, r *http.Request) {
	weekday := time.Now().Weekday()
	eventID := int(weekday)

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/today", getToday).Methods(("GET"))
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
