package helperPrototypes

type VehicleInfo struct {
	Vin                string `json:"vin"`
	CurrentChargeLevel int    `json:"currentChargeLevel"`
	Error              string `json:"error"`
}

func (vehicleInfo *VehicleInfo) HasError() bool {
	return vehicleInfo.Error == "Invalid VIN"
}
