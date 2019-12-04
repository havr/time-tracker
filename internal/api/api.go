package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/havr/time-tracker/internal/models"
	"github.com/havr/time-tracker/internal/stores"
)

func NewRouter(staticPath string, sessionStore *stores.DatabaseSessionStore) http.Handler {
	r := handler{
		sessionStore: sessionStore,
	}

	router := mux.NewRouter()
    router.Methods("GET").Path("/api/v1/sessions").HandlerFunc(r.listSessions)
	router.Methods("POST").Path("/api/v1/sessions").HandlerFunc(r.saveSession)
    router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticPath)))

	return router
}

type handler struct {
	sessionStore *stores.DatabaseSessionStore
}

type saveSessionRequest struct {
	Name     string        `json:"name"`
	Start    time.Time     `json:"startTime"`
	Duration time.Duration `json:"duration"`
}

func (h *handler) saveSession(w http.ResponseWriter, r *http.Request) {
	var request saveSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.sessionStore.SaveSession(&models.Session{
		ID:        uuid.New(),
		Name:      request.Name,
		StartTime: request.Start,
		Duration:  request.Duration,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

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
