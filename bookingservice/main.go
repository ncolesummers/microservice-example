package main

import (
	"flag"

	"github.com/ncolesummers/microservice-example/bookingservice/listener"
	"github.com/ncolesummers/microservice-example/eventsservice/rest"
	"github.com/ncolesummers/microservice-example/lib/configuration"
	msgqueue_amqp "github.com/ncolesummers/microservice-example/lib/msgqueue/amqp"
	"github.com/ncolesummers/microservice-example/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("config", "./configuration/config.json", "path to config file")
	flag.Parse()
	config := configuration.ExtractConfiguration(*confPath)

	dblayer, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	eventListener, err := msgqueue_amqp.NewAMQPEventListener(conn)
	if err != nil {
		panic(err)
	}

	processor := &listener.EventProcessor{eventListener, dblayer}
	go processor.ProcessEvents()

	rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
}
