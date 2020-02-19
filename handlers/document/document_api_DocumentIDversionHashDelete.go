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

// DocumentIDversionHashDelete is the handler for DELETE /document/{documentID}/{versionHash}
// Delete specific version of the document from the knowledge base
func (api DocumentAPI) DocumentIDversionHashDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	versionHash := vars["versionHash"]

	if _, err := db.DeleteOne(context.Background(), bson.M{
		"document.etag": versionHash,
	}); err != nil {
		log.Errorf("unable to delete document: %v", err)
		util.SendUnprocessedEntrityErrorRSP(err, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(204)

}
