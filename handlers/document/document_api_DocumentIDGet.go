// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"encoding/json"
	"github.com/pgonch/knowledge-base/types"
	"net/http"
)

// DocumentIDGet is the handler for GET /document/{documentID}
// Get specific document
func (api DocumentAPI) DocumentIDGet(w http.ResponseWriter, r *http.Request) { // limit := req.FormValue("limit")// offset := req.FormValue("offset")
	w.Header().Set("Content-Type", "application/json")
	var respBody types.DocumentVersionList
	json.NewEncoder(w).Encode(&respBody)
}
