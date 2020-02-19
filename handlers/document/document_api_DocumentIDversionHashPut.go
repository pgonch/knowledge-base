// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"net/http"
	"context"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/pgonch/knowledge-base/util"
	"github.com/pgonch/knowledge-base/types"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

// DocumentIDversionHashPut is the handler for PUT /document/{documentID}/{versionHash}
// Update specific version of the document from the knowledge base
func (api DocumentAPI) DocumentIDversionHashPut(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	documentID := vars["documentID"]
	versionHash := vars["versionHash"]

	var reqBody types.Document

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Debugf("document post: unable to decode: %v", err)
		util.SendUnprocessedEntrityErrorRSP(err, http.StatusUnprocessableEntity, w)
		return
	}

	// transaction
	err := dbClient.UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
	    if err := sessionContext.StartTransaction(); err != nil {
	        return errors.Wrap(err, "unable to start transaction: %v")
	    }

			result := util.DBWrapper{}
			if err := db.FindOne(context.Background(), bson.M{
				"document.document_id":bson.M{
					"$eq": documentID,
				},
				"document.etag":bson.M{
					"$eq": versionHash,
				},
				}).Decode(&result); err != nil {
					sessionContext.AbortTransaction(sessionContext)
			    return errors.Wrap(err,"unable to query collection: %v")
	     	}

				doc, err := util.HashifyDocument(documentID, reqBody,
					reqBody.Version.Created, time.Now().Unix())
				if err != nil {
					sessionContext.AbortTransaction(sessionContext)
			    return errors.Wrap(err, "unable to hashify document: %v")
				}

				result.Document = *doc

				if _, err := db.UpdateOne(context.Background(), bson.M{
					"_id":bson.M{
						"$eq": result.Id,
					},
					}, bson.M{
						"$set":bson.M{
							"document": result.Document,
						},
						}); err != nil {
						sessionContext.AbortTransaction(sessionContext)
				    return errors.Wrap(err, "unable to write to hash document post: %v")
				 	}
    	sessionContext.CommitTransaction(sessionContext)

			reqBody = result.Document
			return nil
	})

	if err != nil {
		log.Error(err)
		util.SendUnprocessedEntrityErrorRSP(err, http.StatusInternalServerError, w)
		return
	}

	json.NewEncoder(w).Encode(&reqBody)

}
