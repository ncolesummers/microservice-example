package daprPub

import (
	"log"
	"os"
	"strings"

	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/ncolesummers/microservice-example/lib/msgqueue"
)

type daprEventListener struct {
	service    daprd.Server
	listenPort string
	pubsubName string
	topic      string
	mapper     msgqueue.EventMapper
}

func NewDaprEventListener() (msgqueue.EventListener, error) {
	var (
		logger         = log.New(os.Stdout, "", 0)
		serviceAddress = getEnvVar("ADDRESS", ":5000")
		pubSubName     = getEnvVar("PUBSUB_NAME", "pubsub")
		topicName      = getEnvVar("TOPIC_NAME", "neworder")
	)

	listener := daprEventListener{
		service:    daprd.Server{},
		listenPort: serviceAddress,
		pubsubName: pubSubName,
		topic:      topicName,
		mapper:     msgqueue.NewEventMapper(),
	}

	return &listener, nil
}

func getEnvVar(key, fallbackValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return strings.TrimSpace(val)
	}
	return fallbackValue
}

func (l *daprEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	//TODO: Finish Implementation
}
