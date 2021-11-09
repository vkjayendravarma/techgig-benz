package helperPrototypes

type DistanceResponse struct {
	Error    string `json:"error"`
	Distance int    `json:"distance"`
}

type ChargingStationsResponse struct {
	ChargingStations []ChargingStation `json:"chargingStations"`
	Error            string            `json:"error"`
}

type ChargingStation struct {
	Name     string `json:"name"`
	Distance int    `json:"distance"`
	Limit    int    `json:"limit"`
}
