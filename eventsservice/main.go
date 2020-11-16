package eventsservice

import (
	"flag"
	"log"

	"github.com/Shopify/sarama"

	"github.com/ncolesummers/microservice-example/eventsservice/rest"
	"github.com/ncolesummers/microservice-example/lib/configuration"
	"github.com/ncolesummers/microservice-example/lib/msgqueue/amqp"
	"github.com/ncolesummers/microservice-example/lib/persistence/dblayer"
	"github.com/streadway/amqp"
)

func main() {
	var eventEmitter msgqueue.eventEmitter

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	log.Println("Connecting to Message Broker...")
	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic(err)

		}
		emitter, err := msqgqueue_amqp.NewAMQPEventEmitter(conn)
		if err != nil {
			panic(err)
		}
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		if err != nil {
			panic(err)
		}
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
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
