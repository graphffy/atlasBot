package structs

type ResponseFromAtlas struct {
	Rides []Ride `json:"rides"`
}

type Ride struct {
	DepartureTime string `json:"departure"`
	Price         int    `json:"onlinePrice"`
	SeatsCount    int    `json:"freeSeats"`
}
