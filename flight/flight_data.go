package flight

type Flight struct {
	FlightDate   string        `json:"flight_date"`
	FlightStatus string        `json:"flight_status"`
	Departure    Departure     `json:"departure"`
	Arrival      Arrival       `json:"arrival"`
	Airline      Airline       `json:"airline"`
	Flight       FlightDetails `json:"flight"`
	Aircraft     Aircraft      `json:"aircraft"`
	Live         interface{}   `json:"live"` // Can be nil or any type
}

type Departure struct {
	Airport         string `json:"airport"`
	Timezone        string `json:"timezone"`
	IATA            string `json:"iata"`
	ICAO            string `json:"icao"`
	Terminal        string `json:"terminal"`
	Gate            string `json:"gate"`
	Delay           int    `json:"delay"`
	Scheduled       string `json:"scheduled"`
	Estimated       string `json:"estimated"`
	Actual          string `json:"actual"`
	EstimatedRunway string `json:"estimated_runway"`
	ActualRunway    string `json:"actual_runway"`
}

type Arrival struct {
	Airport         string `json:"airport"`
	Timezone        string `json:"timezone"`
	IATA            string `json:"iata"`
	ICAO            string `json:"icao"`
	Terminal        string `json:"terminal"`
	Gate            string `json:"gate"`
	Baggage         string `json:"baggage"`
	Delay           int    `json:"delay"`
	Scheduled       string `json:"scheduled"`
	Estimated       string `json:"estimated"`
	Actual          string `json:"actual"`
	EstimatedRunway string `json:"estimated_runway"`
	ActualRunway    string `json:"actual_runway"`
}

type Airline struct {
	Name string `json:"name"`
	IATA string `json:"iata"`
	ICAO string `json:"icao"`
}

type FlightDetails struct {
	Number     string      `json:"number"`
	IATA       string      `json:"iata"`
	ICAO       string      `json:"icao"`
	Codeshared interface{} `json:"codeshared"` // Can be nil
}

type Aircraft struct {
	Registration string `json:"registration"`
	IATA         string `json:"iata"`
	ICAO         string `json:"icao"`
	ICAO24       string `json:"icao24"`
}
