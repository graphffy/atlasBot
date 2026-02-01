package server

import (
	"atlasBot/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var request structs.Request

	fillStructFromRequest(r, &request)

	Server(request, w)

}

func Server(data structs.Request, w http.ResponseWriter) {

	streamResponse := Streaming(w)

	client, url := requestToAtlas(data)

	end := time.After(time.Duration(data.SearchTimeout) * time.Second)

	for {
		select {
		case <-end:
			return
		default:
			filtered := getRidesInfo(client, url, data)
			fmt.Println(filtered)

			fmt.Fprintf(w, "data: ")
			json.NewEncoder(w).Encode(filtered)
			fmt.Fprint(w, "\n\n")

			streamResponse.Flush()

			time.Sleep(time.Duration(data.RequestTimeout) * time.Second)
		}
	}
}

func fillStructFromRequest(r *http.Request, request *structs.Request) {

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		fmt.Println(err)
	}

	request.CityFromId = "c625144"
	request.CityToId = "c625665"
}
