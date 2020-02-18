// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"encoding/json"
	"github.com/pgonch/knowledge-base/types"
	"net/http"
)

// Post is the handler for POST /document
func (api DocumentAPI) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody types.Document

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

}
