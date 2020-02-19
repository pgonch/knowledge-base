// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DocumentAPI is API implementation of /document root endpoint
type DocumentAPI struct {
}

var (
	db       *mongo.Collection
	dbClient *mongo.Client
)

// SetDB sets d as a db
func SetDB(d *mongo.Collection) {
	db = d
}

// SetDBClient sets db client
func SetDBClient(d *mongo.Client) {
	dbClient = d
}
