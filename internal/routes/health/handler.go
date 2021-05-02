package health

import (
	"encoding/json"
	"net/http"
)

// Handler returns health-check
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status: "ok",
	})
}
