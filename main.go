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
	SearchTimeout  int    `json:"searchTimeout"`
	RequestTimeout int    `json:"requestTimeout"`
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

	Client(request, w)

}

func Client(data ClientData, w http.ResponseWriter) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	client := &http.Client{Timeout: 5 * time.Second}

	url := "https://atlasbus.by/api/search" +
		"?from_id=" + data.CityFromId +
		"&to_id=" + data.CityToId +
		"&calendar_width=30" +
		"&date=" + data.Date +
		"&passengers=1"

	end := time.After(time.Duration(data.SearchTimeout) * time.Second)

	for {
		select {
		case <-end:
			return
		default:
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}

			var result SearchResponse
			err = json.NewDecoder(resp.Body).Decode(&result)
			resp.Body.Close()

			if err != nil {
				fmt.Println(err)
				return
			}

			filtered := checker(result, data)

			fmt.Fprintf(w, "data: ")
			json.NewEncoder(w).Encode(filtered)
			fmt.Fprint(w, "\n\n")

			flusher.Flush()

			time.Sleep(time.Duration(data.RequestTimeout) * time.Second)
		}
	}
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

	err := http.ListenAndServe(":9003", nil)
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
