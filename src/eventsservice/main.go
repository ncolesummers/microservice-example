package eventsservice

import (
	"flag"
	"fmt"
	"github.com/ncolesummers/microservice-example/src/eventsservice/rest"
	"github.com/ncolesummers/microservice-example/src/lib/configuration"
	"github.com/ncolesummers/microservice-example/src/lib/persistence/dblayer"
	"log"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	// RESTful API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}

