package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
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
	releaseDate := time.Date(2017, 3, 3, 0, 0, 0, 0, estTimezone)
	daysUntilRelease := math.Floor(releaseDate.Sub(currentDate).Hours() / 24)

	if daysUntilRelease < 0 {
		daysUntilRelease = 0
	}

	releaseString := fmt.Sprintf("There are *%v days* until the Nintendo Switch is released.", daysUntilRelease)

	m := &MessageResponse{
		Text:         releaseString,
		ResponseType: "in_channel",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
