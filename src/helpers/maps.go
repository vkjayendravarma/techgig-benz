package helpers

import (
	"encoding/json"
	"sync"

	"github.com/vkjayendravarma/techgig-benz-v2/src/gatewaysAndApis/techgig"
	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/controllerResponses"
	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/helperPrototypes"
)

/*
Helper for EvRoutePlan. Gets distance from source to destination and updates in the response

@query: controller response pointer
*/
func DistanceFinder(query *controllerResponses.EvRoutePlanResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	responseBodyByteArray := techgig.PostRequest(`{ "source": "`+query.Source+`","destination": "`+query.Destination+`"  }`, "distance")

	if responseBodyByteArray == nil {
		// add error if api call has error
		query.AddError(9999, "Technical Error")
		return
	}

	var distanceStats *helperPrototypes.DistanceResponse

	json.Unmarshal(responseBodyByteArray, &distanceStats)

	if distanceStats.Error != "" {
		// add error if api response has error
		query.AddError(9999, "Technical Error"+distanceStats.Error)
		return
	}
	// update distance in controller responce
	query.Distance = distanceStats.Distance

}

/*
Helper for EvRoutePlan. Get all charging stations on the way

@query: controller response pointer

returns array of charging stations
*/
func ChargingStationsFinder(query *controllerResponses.EvRoutePlanResponse) []helperPrototypes.ChargingStation {

	responseBodyByteArray := techgig.PostRequest(`{ "source": "`+query.Source+`","destination": "`+query.Destination+`"  }`, "charging_stations")

	if responseBodyByteArray == nil {
		// add error if api call has error
		query.AddError(9999, "Technical Error")
		return nil
	}

	var chargingStationsResponse helperPrototypes.ChargingStationsResponse
	// convert response byte array to struct
	json.Unmarshal(responseBodyByteArray, &chargingStationsResponse)

	if chargingStationsResponse.Error != "" {
		// add error if api response has error
		query.AddError(9999, chargingStationsResponse.Error)
		return nil
	}

	// else return charging stations
	return chargingStationsResponse.ChargingStations
}
