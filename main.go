package microservice_example

import (
	"github.com/gorilla/mux"
	"net/http"
)

func ServeAPI(endpoint string) error {
	handler := &eventservicehandler{}
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}

type eventServiceHandler struct{}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {}

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request) {}

func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {}
