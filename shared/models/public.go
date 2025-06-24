package models

type Performance struct {
	Title      string  `json:"title"`
	Venue      string  `json:"venue"`
	ExactDate  string  `json:"exactDate"`
	Date       string  `json:"date"`
	TicketLink string  `json:"ticketLink"`
	Programme  []Piece `json:"programme"`
}

type Piece struct {
	Composer string `json:"composer"`
	Title    string `json:"title"`
}

type Video struct {
	Title         string `json:"title"`
	ExtendedTitle string `json:"extendedTitle"`
	EmbedURL      string `json:"embedURL"`
}

type ContactDetails struct {
	Position string `json:"position"`
	Location string `json:"location"`
	TelNum   string `json:"telNum"`
	TelText  string `json:"telText"`
	Email    string `json:"email"`
}
