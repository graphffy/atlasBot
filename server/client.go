package server

import (
	"atlasBot/structs"
	"net/http"
	"time"
)

func requestToAtlas(data structs.Request) (*http.Client, string) {

	client := &http.Client{Timeout: 5 * time.Second}

	url := "https://atlasbus.by/api/search" +
		"?from_id=" + data.CityFromId +
		"&to_id=" + data.CityToId +
		"&calendar_width=30" +
		"&date=" + data.Date +
		"&passengers=1"

	return client, url
}
