// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"encoding/json"
	"github.com/pgonch/knowledge-base/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	log "github.com/sirupsen/logrus"
	"context"
	"net/http"
)

// DocumentIDversionHashGet is the handler for GET /document/{documentID}/{versionHash}
// Get specific version of the document
func (api DocumentAPI) DocumentIDversionHashGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	documentID := vars["documentID"]
	versionHash := vars["versionHash"]

	result := util.DBWrapper{}
	if err := db.FindOne(context.Background(), bson.M{
		"document.document_id":bson.M{
			"$eq": documentID,
		},
		"document.etag":bson.M{
			"$eq": versionHash,
		},
		}).Decode(&result); err != nil {
			log.Errorf("unable to decode query result: %v", err )
		}

	json.NewEncoder(w).Encode(&result.Document)
}
