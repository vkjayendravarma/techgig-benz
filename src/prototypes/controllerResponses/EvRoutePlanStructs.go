package controllerResponses

import (
	"strconv"
	"time"

	"github.com/vkjayendravarma/techgig-benz-v2/src/prototypes/helperPrototypes"
)

type EvRoutePlanResponse struct {
	TransactionId        string                             `json:"transactionId"`
	Vin                  string                             `json:"vin"`
	Source               string                             `json:"source"`
	Destination          string                             `json:"destination"`
	Distance             int                                `json:"distance,omitempty"`
	CurrentChargingLevel int                                `json:"currentChargeLevel,omitempty"`
	ChargingStations     []helperPrototypes.ChargingStation `json:"chargingStations,omitempty"`
	Error                []EvRoutePlanErrorResponse         `json:"error,omitempty"`
	IsChargingRequired   bool                               `json:"isChargingRequired"`
}

type EvRoutePlanErrorResponse struct {
	Id          uint16 `json:"id"`
	Description string `json:"description"`
}

func (routePlan *EvRoutePlanResponse) Init() {
	timeStamp := time.Now().UTC()
	routePlan.TransactionId = strconv.Itoa(timeStamp.Year()) + strconv.Itoa(timeStamp.YearDay()) + strconv.Itoa(timeStamp.Hour()) + strconv.Itoa(timeStamp.Minute()) + strconv.Itoa(timeStamp.Second()) + strconv.Itoa(int(timeStamp.UnixMilli()))
}

func (routePlan *EvRoutePlanResponse) Invalid() bool {
	return routePlan.Vin == "" || routePlan.Source == "" || routePlan.Destination == ""
}

func (routePlan *EvRoutePlanResponse) AddError(code uint16, description string) {
	newError := EvRoutePlanErrorResponse{code, description}
	routePlan.Error = append(routePlan.Error, newError)
}
