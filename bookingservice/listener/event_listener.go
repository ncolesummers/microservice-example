package listener

import (
	"todo.com/myevents/contracts"
	"log"
	"github.com/ncolesummers/microservice-example/lib/msgqueue"
	"github.com/ncolesummers/microservice-example/lib/persistence"
	"gopkg.in/mgo.v2/bson"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database	persistence.DatabaseHandler
}

func (p *EventProcessor) ProcessEvents() error {
	log.Println("Listening to events...")

	received, errors, err := p.EventListener.Listen("eventCreated")
	if err != nil {
		return err
	}

	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("received error while processing msg: %s", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s", e.ID, e)
		p.Database.AddLocation(persistence.Location{ID: bson.ObjectId(e.ID)})
	default:
		log.Printf("unknown event: %t", e)
	}
}