// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"net/http"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/pgonch/knowledge-base/util"

	log "github.com/sirupsen/logrus"

)

// DocumentIDDelete is the handler for DELETE /document/{documentID}
// Delete document with documentID from the knowledge base
func (api DocumentAPI) DocumentIDDelete(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		documentID := vars["documentID"]

		if _, err := db.DeleteMany(context.Background(), bson.M{
			"document.document_id": documentID,
		}); err != nil {
			log.Errorf("unable to delete document: %v", err)
			util.SendUnprocessedEntrityErrorRSP(err, http.StatusInternalServerError, w)
			return
		}

		w.WriteHeader(204)
}
