package controller

import (
	"cache-service/common"
	"encoding/json"
	"net/http"
)

type HealthCheck struct{}

var Ping HealthCheck

//Ping for healthcheck
func (h *HealthCheck) Ping(w http.ResponseWriter, r *http.Request) {
	response := common.ResponseMapper("", http.StatusOK, "OK")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}
