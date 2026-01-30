package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ClientData struct {
	Date           string `json:"date"`
	TimeFrom       int    `json:"timeFrom"`
	TimeTo         int    `json:"timeTo"`
	CityFrom       string `json:"cityFrom"`
	CityTo         string `json:"cityTo"`
	SearchTimeout  string `json:"searchTimeout"`
	RequestTimeout string `json:"requestTimeout"`
	CityFromId     string
	CityToId       string
}

type SearchResponse struct {
	Trips []Trip `json:"rides"`
}

type Trip struct {
	Time  string `json:"departure"`
	Price int    `json:"onlinePrice"`
	Seats int    `json:"freeSeats"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var request ClientData
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(err)
	}
	fmt.Println(request)

	request.CityFromId = "c625144"
	request.CityToId = "c625665"

	go Client(request)

}

func Client(data ClientData) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	src := "https://atlasbus.by/api/search?from_id=" + data.CityFromId + "&to_id=" + data.CityToId +
		"&calendar_width=30&date=" + data.Date + "&passengers=1"

	page, err := client.Get(src)
	if err != nil {
		fmt.Println(err)
	}
	defer page.Body.Close()

	var result SearchResponse
	if err := json.NewDecoder(page.Body).Decode(&result); err != nil {
		fmt.Println(err)
	}

	r := checker(result, data)

	fmt.Println(r)

}

func checker(s SearchResponse, d ClientData) SearchResponse {
	var result SearchResponse
	for _, v := range s.Trips {
		if v.Seats != 0 {
			houres, _ := strconv.Atoi(v.Time[11:13])
			if houres >= d.TimeFrom && houres < d.TimeTo {
				result.Trips = append(result.Trips, v)
			}
		}

	}
	return result
}

func main() {
	http.HandleFunc("/", Handler)

	err := http.ListenAndServe(":9008", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
{
    "date" : "26.01",
    "timeFrom" : "16.00",
    "timeTo" : "19.00",
    "cityFrom" : "Минск",
    "CityTo" : "Могилев",
    "SearchTimeout" : "10min",
    "RequestTimeout" : "1min"
}
*/
//https://atlasbus.by/Маршруты/Минск/Могилев?date=2026-02-03&passengers=1&from=c625144&to=c625665
//https://atlasbus.by/api/search?from_id=c625144&to_id=c625665&calendar_width=30&date=2026-02-01&passengers=1&operatorId=
