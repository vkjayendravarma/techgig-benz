package controllers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/vkjayendravarma/techgig-benz-v2/src/helpers"
	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/controllerResponses"
)

/*
@type: POST
@route: /api/plantrip
@desc: controller to check wether trip is possible and plan with optimal charging staions
@access: PUBLIC
*/
func EvRoutePlan(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var response controllerResponses.EvRoutePlanResponse
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		json.NewEncoder(w).Encode("Error missing or invalid request format")
		return
	}
	if response.Invalid() {
		json.NewEncoder(w).Encode("Error missing or invalid request format")
		return
	}

	// initiate transaction
	response.Init()

	var wg sync.WaitGroup
	wg.Add(2)

	// get vehicle battery status
	go helpers.VehicleBatteryStatus(&response, &wg)

	// get distance
	go helpers.DistanceFinder(&response, &wg)

	wg.Wait()

	if len(response.Error) > 0 {
		json.NewEncoder(w).Encode(response)
		return
	}

	if response.Distance > response.CurrentChargingLevel {
		// if charging required
		response.IsChargingRequired = true

		// charging stations on route
		chargingStations := helpers.ChargingStationsFinder(&response)

		// add optimal charging stops
		helpers.OptimizePitStops(&response, chargingStations)
	}

	json.NewEncoder(w).Encode(response)
}
