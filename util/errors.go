package util

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/pgonch/knowledge-base/types"
)

// SendUnprocessedEntrityErrorRSP sends typical error RSP back
func SendUnprocessedEntrityErrorRSP(err error, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	log.Error(err)
	if e := json.NewEncoder(w).Encode(&types.ErrorDescription{
		Error_description: err.Error(),
	}); e != nil {
		log.Errorf("Problem with encoding error: %v", e)
	}
}
