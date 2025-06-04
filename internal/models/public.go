package models

type Event struct {
	Title     string  `json:"title"`
	Venue     string  `json:"venue"`
	Programme []Piece `json:"programme"`
}

type Piece struct {
	Composer string `json:"composer"`
	Title    string `json:"title"`
}

