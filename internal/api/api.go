package api

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/havr/time-tracker/internal/stores"
)

type handler struct {
	sessionStore stores.SessionStore
}

func NewRouter(staticPath string, sessionStore stores.SessionStore) http.Handler {
	r := handler{
		sessionStore: sessionStore,
	}

	router := mux.NewRouter()
	router.Methods("GET").Path("/api/v1/sessions").HandlerFunc(r.listSessions)
	router.Methods("POST").Path("/api/v1/sessions").HandlerFunc(r.saveSession)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticPath)))

	return router
}
