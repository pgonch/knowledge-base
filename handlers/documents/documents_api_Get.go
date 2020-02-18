// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package documents

import (
	"encoding/json"
	"github.com/pgonch/knowledge-base/types"
	"net/http"
)

// Get is the handler for GET /documents
// Get all documents stored in the knowledge base
func (api DocumentsAPI) Get(w http.ResponseWriter, r *http.Request) { // limit := req.FormValue("limit")// offset := req.FormValue("offset")
	w.Header().Set("Content-Type", "application/json")
	var respBody types.DocumentsList
	json.NewEncoder(w).Encode(&respBody)
}
