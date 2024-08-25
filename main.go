package main

import (
	"flightTracker/flight"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting flight tracker server ...")
	http.HandleFunc("/flights/", flight.HandleFlightRequest)
	http.ListenAndServe(":8080", nil)
}
