package helpers

import (
	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/controllerResponses"
	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/helperPrototypes"
)

/*
	Find optimal charging station to topup which are crossed

	@query: controller response pointer

	@stations: charging stations with are crossed

	@rangeCanBeCovered: intial range + past aquired range during trip

	@numberOfStations: len(stations)

	This function removes used charging station from available sharging stations adds it to controller response and returns charging limit which can be used
*/
func aquireRange(query *controllerResponses.EvRoutePlanResponse, stations []helperPrototypes.ChargingStation, rangeCanBeCovered int, numberOfStations int) (int, []helperPrototypes.ChargingStation) {
	maxRangeStation := stations[0]
	maxRangeIndex := 0
	maxRange := rangeCanBeCovered - maxRangeStation.Distance + maxRangeStation.Limit

	if numberOfStations > 1 {
		index := 1

		for index < numberOfStations {
			currentAquireingRange := rangeCanBeCovered - stations[index].Distance + stations[index].Limit
			if currentAquireingRange > maxRange {
				maxRangeIndex = index
				maxRangeStation = stations[index]
				maxRange = currentAquireingRange
			}

			index++
		}
	}

	// add used charging station to controller response
	query.ChargingStations = append(query.ChargingStations, maxRangeStation)

	// update next usable charging stations
	stations = append(stations[:maxRangeIndex], stations[maxRangeIndex+1:]...)

	return maxRangeStation.Limit, stations
}

/*
Optimal charging stops driver function

@query: controller response pointer
*/
func OptimizePitStops(query *controllerResponses.EvRoutePlanResponse, stations []helperPrototypes.ChargingStation) {

	// add destination as a charging station with zero limit
	stations = append(stations, helperPrototypes.ChargingStation{"destination", query.Distance, 0})

	var stationsCrossed []helperPrototypes.ChargingStation

	distanceToCover := query.Distance
	batteryStatus := query.CurrentChargingLevel
	rangeCanBeCovered := query.CurrentChargingLevel
	prev := 0
	stationNumber := 0

	stationsLen := len(stations)

	for stationNumber < stationsLen {

		currentBatteryStatus := batteryStatus - (stations[stationNumber].Distance - prev)

		if currentBatteryStatus <= 0 {
			if stations[stationNumber].Distance == rangeCanBeCovered {
				stationsCrossed = append(stationsCrossed, stations[stationNumber])
			}
			numberOfStationsCrossed := len(stationsCrossed)
			if numberOfStationsCrossed == 0 {
				// if no charging staions crossed previously
				query.AddError(9999, "Unable to reach the destination with the current charge level")
				query.ChargingStations = nil
				return
			}
			var aquiredRange int
			aquiredRange, stationsCrossed = aquireRange(query, stationsCrossed, rangeCanBeCovered, numberOfStationsCrossed)
			batteryStatus = batteryStatus + aquiredRange
			if batteryStatus > 100 {
				batteryStatus = 100
			}
			rangeCanBeCovered += batteryStatus
			if rangeCanBeCovered >= distanceToCover {
				break
			}
			batteryStatus -= (stations[stationNumber].Distance - prev)
			if stations[stationNumber].Distance < rangeCanBeCovered {
				stationsCrossed = append(stationsCrossed, stations[stationNumber])
			}
		} else {
			batteryStatus = currentBatteryStatus
			stationsCrossed = append(stationsCrossed, stations[stationNumber])
		}
		prev = stations[stationNumber].Distance - prev
		stationNumber++

	}

}
