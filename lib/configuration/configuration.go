package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ncolesummers/microservice-example/lib/persistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
	RestfulTLSEPDefault = "localhost:9191"
	MessageBrokerTypeDefault = "amqp"
	AMQPMessageBrokerDefault = "amqp://guest:guest@localhost:5672"
	KafkaMessageBrokerDefault = []string{"localhost:9092"}
)

type ServiceConfig struct {
	Databasetype      dblayer.DBTYPE `json:"databasetype"`
	DBConnection      string         `json:"dbconnection"`
	RestfulEndpoint   string         `json:"restfulapi_endpoint"`
	RestfulTLSEndPint string         `json:"restfulapi-tlsendpoint"`
	MessageBrokerType   string         `json:"message_broker_type"`
	AMQPMessageBroker   string         `json:"amqp_message_broker"`
	KafkaMessageBrokers []string       `json:"kafka_message_brokers"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		RestfulTLSEPDefault,
		MessageBrokerTypeDefault,
		AMQPMessageBrokerDefault
		KafkaMessageBrokerDefault
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found.  Continuing with default values.")
		// return conf, err
	}

	json.NewDecoder(file).Decode(&conf)

		if v := os.Getenv("LISTEN_URL"); v != "" {
		conf.RestfulEndpoint = v
	}

	if v := os.Getenv("MONGO_URL"); v != "" {
		conf.Databasetype = "mongodb"
		conf.DBConnection = v
	}

	if v := os.Getenv("AMQP_BROKER_URL"); v != "" {
		conf.MessageBrokerType = "amqp"
		conf.AMQPMessageBroker = v
	} else if v := os.Getenv("KAFKA_BROKER_URLS"); v != "" {
		conf.MessageBrokerType = "kafka"
		conf.KafkaMessageBrokers = strings.Split(v, ",")
	}

	return conf, nil
}