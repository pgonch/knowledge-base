// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"encoding/json"
	"net/http"
	"time"
	"context"

	"github.com/google/uuid"
	"github.com/pgonch/knowledge-base/types"
	"github.com/pgonch/knowledge-base/util"

	log "github.com/sirupsen/logrus"
)

// Post is the handler for POST /document
func (api DocumentAPI) Post(w http.ResponseWriter, r *http.Request) {
	createDocument(uuid.New().String(), w, r)
}


func createDocument(id string, w http.ResponseWriter, r *http.Request) {
	var reqBody types.Document

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Debugf("document post: unable to decode: %v", err)
		util.SendUnprocessedEntrityErrorRSP(err, http.StatusUnprocessableEntity, w)
		return
	}

	doc, err  := util.HashifyDocument(id, reqBody, time.Now().Unix(), 0)
	if err != nil {
		log.Error(err)
		util.SendUnprocessedEntrityErrorRSP(err, http.StatusInternalServerError, w)
		return
	}

	var dbWrapper util.DBWrapper

	dbWrapper.Id = uuid.New().String()
	dbWrapper.Document = *doc

	if _, err := db.InsertOne(context.Background(), dbWrapper); err != nil {
		log.Errorf("unable to write to db: document post: %v", err)
		util.SendUnprocessedEntrityErrorRSP(err, http.StatusUnprocessableEntity, w)
		return
	}

	json.NewEncoder(w).Encode(&reqBody)
}
