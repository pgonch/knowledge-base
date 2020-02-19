// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package document

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DocumentIDPost is the handler for POST /document/{documentID}
// Create specific version of the document from the knowledge base
func (api DocumentAPI) DocumentIDPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	documentID := vars["documentID"]

	createDocument(documentID, w, r)
}
