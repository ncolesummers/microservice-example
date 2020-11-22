package main

import (
	"flag"
	"log"

	"github.com/Shopify/sarama"

	"github.com/ncolesummers/microservice-example/eventsservice/rest"
	"github.com/ncolesummers/microservice-example/lib/configuration"
	"github.com/ncolesummers/microservice-example/lib/msgqueue"
	msqgqueue_amqp "github.com/ncolesummers/microservice-example/lib/msgqueue/amqp"
	"github.com/ncolesummers/microservice-example/lib/msgqueue/kafka"
	"github.com/ncolesummers/microservice-example/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", `./configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, err := configuration.ExtractConfiguration(*confPath)
	if err != nil {
		panic(err)
	}

	log.Println("Connecting to Message Broker...")
	// log.Println(config.MessageBrokerType)
	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic(err)

		}

		eventEmitter, err = msqgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}
	case "kafka":
		// log.Println(config.KafkaMessageBrokers)
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		if err != nil {
			log.Fatal(err)
		}

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	log.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful... ")

	log.Println("Serving API...")
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPoint, dbhandler, eventEmitter)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
