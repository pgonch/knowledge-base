// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package documents

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DocumentsAPI is API implementation of /documents root endpoint
type DocumentsAPI struct {
}

var (
	db *mongo.Collection
)

// SetDB sets d as a db
func SetDB(d *mongo.Collection) {
	db = d
}
