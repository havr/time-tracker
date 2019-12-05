package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/havr/time-tracker/internal/models"
)

func (h *handler) listSessions(w http.ResponseWriter, r *http.Request) {
	var timeRange time.Duration

	rangeParam := r.URL.Query().Get("range")
	switch rangeParam {
	case "day":
		timeRange = time.Hour * 24
	case "week":
		timeRange = time.Hour * 24 * 7
	case "month":
		timeRange = time.Hour * 24 * 30
	default:
		http.Error(w, fmt.Sprintf("invalid time range: %s", rangeParam), http.StatusBadRequest)
		return
	}

	result, err := h.sessionStore.ListSessions(time.Now().Add(-1 * timeRange))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if result == nil {
		result = []models.Session{}
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
