// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"net/http"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/pgonch/knowledge-base/util"
	"github.com/pgonch/knowledge-base/types"

	log "github.com/sirupsen/logrus"

)

// DocumentIDGet is the handler for GET /document/{documentID}
// Get specific document
func (api DocumentAPI) DocumentIDGet(w http.ResponseWriter, r *http.Request) { // limit := req.FormValue("limit")// offset := req.FormValue("offset")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	documentID := vars["documentID"]


	cur, err := db.Find(context.Background(), bson.M{
		"document.document_id":bson.M{
			"$eq": documentID,
		},
	})

	if err != nil {
		log.Errorf("unable to find record from db to delete: %v", err )
		util.SendUnprocessedEntrityErrorRSP(err,
			http.StatusInternalServerError, w)
		return
	}

	var respBody types.DocumentVersionList


	for cur.Next(context.Background()) {

		var elem util.DBWrapper
    err := cur.Decode(&elem)
    if err != nil {
        log.Errorf("unable to decode resp from db: %v", err )
				util.SendUnprocessedEntrityErrorRSP(err,
					http.StatusInternalServerError, w)
				return
    }

		respBody.Items = append(respBody.Items, elem.Document.Version)
	}

	json.NewEncoder(w).Encode(&respBody)
}
