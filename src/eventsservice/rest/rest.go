package rest

import (
	"github.com/gorilla/mux"
	"github.com/ncolesummers/microservice-example/src/lib/persistence"
	"net/http"
)

func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(dbHandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}