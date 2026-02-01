package server

import (
	"atlasBot/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func getRidesInfo(client *http.Client, url string, data structs.Request) structs.ResponseFromAtlas {

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var result structs.ResponseFromAtlas
	err = json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	filtered := filterRidesInfo(result, data)

	return filtered
}

func filterRidesInfo(s structs.ResponseFromAtlas, d structs.Request) structs.ResponseFromAtlas {
	var result structs.ResponseFromAtlas
	for _, v := range s.Rides {
		if v.SeatsCount != 0 {
			houres, _ := strconv.Atoi(v.DepartureTime[11:13])
			if houres >= d.TimeFrom && houres < d.TimeTo {
				result.Rides = append(result.Rides, v)
			}
		}

	}
	return result
}
