package builder

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/ncolesummers/microservice-example/lib/msgqueue"
	"github.com/ncolesummers/microservice-example/lib/msgqueue/amqp"
	"github.com/ncolesummers/microservice-example/lib/msgqueue/kafka"
	dapr "github.com/dapr/go-sdk/client"
)

const (
	daprPort = "3500"
)

var (
	port string
)

func NewEventListenerFromEnvironment() (msgqueue.EventListener, error) {
	var listener msgqueue.EventListener
	var err error

	if url := os.Getenv("AMQP_URL"); url != "" {
		log.Printf("connecting to AMQP broker at %s", url)

		listener, err = amqp.NewAMQPEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		}
	} else if port = os.Getenv("DAPR_GRPC_PORT"); len(port) == 0 {
		port = daprPort
		log.Printf("Using dapr for pubsub at %s", port)

		client, err := dapr.NewClientWithPort(port)
		if err != nil {
			panic(err)
		}
		defer client.Close()
		ctx := context.Background()

	} else if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		log.Printf("connecting to Kafka brokers at %s", brokers)

		listener, err = kafka.NewKafkaEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("No broker specified")
	}

	return listener, nil
}
