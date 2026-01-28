package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ClientData struct {
	Date           string `json:"date"`
	TimeFrom       string `json:"timeFrom"`
	TimeTo         string `json:"timeTo"`
	CityFrom       string `json:"cityFrom"`
	CityTo         string `json:"cityTo"`
	SearchTimeout  string `json:"searchTimeout"`
	RequestTimeout string `json:"requestTimeout"`
	CityFromId     string
	CityToId       string
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
	src := "https://atlasbus.by/Маршруты/" + data.CityFrom + "/" + data.CityTo + "?date=" +
		data.Date + "&passengers=1&from=" + data.CityFromId + "&to=" + data.CityToId
	fmt.Println(src)
	//client := http.Client{}

}

func main() {
	http.HandleFunc("/", Handler)

	err := http.ListenAndServe(":9090", nil)
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
