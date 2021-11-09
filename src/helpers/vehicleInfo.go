package helpers

import (
	"encoding/json"
	"sync"

	"github.com/vkjayendravarma/techgig-benz-v2/src/gatewaysAndApis/techgig"
	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/controllerResponses"
	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/helperPrototypes"
)

func VehicleBatteryStatus(query *controllerResponses.EvRoutePlanResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	responseBodyByteArray := techgig.PostRequest(`{ "vin": "`+query.Vin+`" }`, "charge_level")

	if responseBodyByteArray == nil {
		// add error if api call has error
		query.AddError(9999, "Technical Error")
		return
	}

	// convert response byte array to struct
	var vehicleStats *helperPrototypes.VehicleInfo
	json.Unmarshal(responseBodyByteArray, &vehicleStats)

	if vehicleStats.Error != "" {
		// add error if api response has error
		query.AddError(9999, "Technical Error"+vehicleStats.Error)
		return
	}

	// else update battery status to response
	query.CurrentChargingLevel = vehicleStats.CurrentChargeLevel

}
