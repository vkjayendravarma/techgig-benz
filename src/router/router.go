package router

import (
	"github.com/vkjayendravarma/techgig-benz-v2/src/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/healthcheck", controllers.HealthChecker).Methods("GET")
	router.HandleFunc("/api/plantrip", controllers.EvRoutePlan).Methods("POST")

	return router
}
