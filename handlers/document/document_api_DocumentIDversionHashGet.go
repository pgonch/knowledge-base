// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"encoding/json"
	"github.com/pgonch/knowledge-base/types"
	"net/http"
)

// DocumentIDversionHashGet is the handler for GET /document/{documentID}/{versionHash}
// Get specific version of the document
func (api DocumentAPI) DocumentIDversionHashGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var respBody types.Document
	json.NewEncoder(w).Encode(&respBody)
}
