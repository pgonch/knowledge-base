// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package documents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"

	"github.com/pgonch/knowledge-base/types"
	"github.com/pgonch/knowledge-base/util"


	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

type Wrap struct {
	Id string `bson:"_id"`
	types.Document
}

// Get is the handler for GET /documents
// Get all documents stored in the knowledge base
func (api DocumentsAPI) Get(w http.ResponseWriter, r *http.Request) { // limit := req.FormValue("limit")// offset := req.FormValue("offset")
	w.Header().Set("Content-Type", "application/json")

	// limit, err := util.PaginationLimitParseAndValidate(r.FormValue("limit"))
	// if err != nil {
	// 	log.Debugf("wrong limit is provided: %v", err)
	// 	util.SendUnprocessedEntrityErrorRSP(err, http.StatusUnprocessableEntity, w)
	// 	return
	// }
	//
	// offset, err := util.PaginationOffsetParseAndValidate(r.FormValue("offset"))
	// if err != nil {
	// 	log.Debugf("wrong offset is provided: %v", err)
	// 	util.SendUnprocessedEntrityErrorRSP(err, http.StatusUnprocessableEntity, w)
	// 	return
	// }
	//
	// var opts options.FindOptions
	// if offset != 0 {
	// 	opts.SetSkip(int64(offset))
	// }
	//
	// if limit != 0 {
	// 	opts.SetLimit(int64(limit))
	// }

	values, err := db.Distinct(context.Background(), "document.document_id",
	bson.D{}, options.Distinct())
	if err != nil {
		log.Errorf("unable to query db")
		util.SendUnprocessedEntrityErrorRSP(fmt.Errorf("something went wrong"),
		 http.StatusInternalServerError, w)
	}

	result := types.DocumentsList{}
	for _, val := range values  {
		result.Items = append(result.Items, fmt.Sprintf("%s", val))
  }

	json.NewEncoder(w).Encode(&result)
}
