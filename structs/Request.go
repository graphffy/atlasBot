package structs

type Request struct {
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
