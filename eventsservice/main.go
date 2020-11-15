package eventsservice

import (
	"flag"
	"log"

	"github.com/ncolesummers/microservice-example/eventsservice/rest"
	"github.com/ncolesummers/microservice-example/lib/configuration"
	"github.com/ncolesummers/microservice-example/lib/persistence/dblayer"
	"github.com/streadway/amqp"
	"todo.com/myevents/lib/msgqueue/amqp"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config := configuration.ExtractConfiguration(*confPath)
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	emitter, err := msqgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	log.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful... ")
	//RESTful API start
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
