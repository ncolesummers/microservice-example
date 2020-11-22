package main

import (
	"flag"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/ncolesummers/microservice-example/bookingservice/listener"
	"github.com/ncolesummers/microservice-example/bookingservice/rest"
	"github.com/ncolesummers/microservice-example/lib/configuration"
	"github.com/ncolesummers/microservice-example/lib/msgqueue"
	msgqueue_amqp "github.com/ncolesummers/microservice-example/lib/msgqueue/amqp"
	"github.com/ncolesummers/microservice-example/lib/msgqueue/kafka"
	"github.com/ncolesummers/microservice-example/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		eventListener msgqueue.EventListener
		eventEmitter  msgqueue.EventEmitter
	)
	// fmt.Println("starting")
	confPath := flag.String("config", "./configuration/config.json", "path to config file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confPath)
	// panicIfErr(err)

	fmt.Println(config.RestfulEndpoint)
	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		panicIfErr(err)

		eventListener, err = msgqueue_amqp.NewAMQPEventListener(conn, "events", "booking")
		panicIfErr(err)

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		panicIfErr(err)
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		panicIfErr(err)

		eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
		panicIfErr(err)

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		panicIfErr(err)
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	processor := listener.EventProcessor{eventListener, dbhandler}
	go processor.ProcessEvents()

	rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
}
