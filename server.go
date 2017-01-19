package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type MessageResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func SwitchHandler(w http.ResponseWriter, r *http.Request) {
	estTimezone, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("Error thrown from timezone: %v", err)
	}

	currentDate := time.Now()
	releaseDate := time.Date(2017, 5, 3, 0, 0, 0, 0, estTimezone)
	daysUntilRelease := releaseDate.Sub(currentDate).Hours() / 24
	log.Printf("Days Until Nintendo Switch Release: %v", daysUntilRelease)
	log.Print("Posting to this body")
	w.WriteHeader(http.StatusOK)
	releaseString := fmt.Sprintf("There are *%v days* until the Nintendo Switch is released.", daysUntilRelease)

	m := &MessageResponse{
		Text:         releaseString,
		ResponseType: "in_channel",
	}

	/*response, err := json.Marshal(m)
	if err != nil {
		log.Printf("Something went wrong! %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	log.Print(m)
	//w.Write(response)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/switch", SwitchHandler)
	r.HandleFunc("/", HomeHandler)
	log.Print("Starting up the server.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
