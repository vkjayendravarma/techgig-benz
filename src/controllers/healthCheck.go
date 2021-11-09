package controllers

import "net/http"

/*
@type: GET
@route: /api/healthcheck
@desc: Api health check route
@access: PUBLIC
*/
func HealthChecker(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api is working fine"))
}
