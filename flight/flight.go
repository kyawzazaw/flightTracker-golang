package flight

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleFlightRequest(w http.ResponseWriter, r *http.Request) {
	flightNumber := r.URL.Path[len("/flights/"):]

	apiKey := "2a0d3f2c3b2f73d9c1d8ee8824905e30"
	apiURL := fmt.Sprintf("https://api.aviationstack.com/v1/flights?flight_iata=%s&access_key=%s", flightNumber, apiKey)

	fmt.Println("Making API request to: %s", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching flight data: %v", err)
		http.Error(w, "Error fetching flight data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the raw response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: %v", err)
		http.Error(w, "Error fetching flight data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close() // Ensure the body is closed even after reading

	// Print the raw response body
	fmt.Println("Raw API response: %s", string(bodyBytes))

	// Continue with decoding and handling the response as before
	fmt.Println("Received API response with status code: %d", resp.StatusCode)

	var flightData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&flightData)
	if err != nil {
		fmt.Println("Error decoding flight data: %v", err)
		http.Error(w, "Error decoding flight data", http.StatusInternalServerError)
		return
	}

	flights, ok := flightData["data"].([]interface{})
	if !ok {
		fmt.Println("Error: 'data' key is not an array")
		http.Error(w, "Error fetching flight data", http.StatusInternalServerError)
		return
	}

	for _, flight := range flights {
		flightMap := flight.(map[string]interface{})
		flightDate, _ := flightMap["flight_date"].(string)
		flightStatus, _ := flightMap["flight_status"].(string)
		// ... and so on for other fields
		fmt.Printf("Flight Date: %s, Status: %s\n", flightDate, flightStatus)
		// Process or display flight information
	}

	// Encode the processed flight data or a custom response structure as needed
	// json.NewEncoder(w).Encode(processedFlightData)
}
