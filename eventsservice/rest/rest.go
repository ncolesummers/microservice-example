package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ncolesummers/microservice-example/lib/msgqueue"
	"github.com/ncolesummers/microservice-example/lib/persistence"
)

func ServeAPI(endpoint, tlsendpoint string, databaseHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
	handler := NewEventHandler(databaseHandler, eventEmitter)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)
	go func() { httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "cert.pem", "key.pem", r) }()
	go func() { httpErrChan <- http.ListenAndServe(endpoint, r) }()
	return httpErrChan, httptlsErrChan
}
