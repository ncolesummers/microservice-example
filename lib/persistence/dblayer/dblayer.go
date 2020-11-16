package dblayer

import (
	"github.com/ncolesummers/microservice-example/lib/persistence"
	"github.com/ncolesummers/microservice-example/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB    DBTYPE = "mongodb"
	DYNAMODB   DBTYPE = "dynamodb"
	DOCUMENTDB DBTYPE = "documentdb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
