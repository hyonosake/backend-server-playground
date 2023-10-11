package healthcheck

import (
	"encoding/json"
	"net/http"

	"github.com/hyonosake/backend-server-playground/internal/rest_handlers/healthcheck/entity"
)

const HealthcheckRequestRoute = "/internal/healthcheck"

func Handle() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := entity.HealthcheckResponse{Alive: true}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}
}

func NewHealthcheckHandler() func(http.ResponseWriter, *http.Request) {
	return Handle()
}
