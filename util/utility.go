package util

import (
	"crypto/md5"

	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/pgonch/knowledge-base/types"
	log "github.com/sirupsen/logrus"
)

// ParseLogLevel parses string to logrus level object
func ParseLogLevel(ll string) error {

	lvl, err := log.ParseLevel(ll)
	if err != nil {
		return err
	}

	if lvl == log.DebugLevel || lvl == log.TraceLevel {
		log.SetReportCaller(true)
	}

	log.SetLevel(lvl)

	return nil
}

const (
	// QueryWarningLimitSec is a warning toleration limit for query duration
	QueryWarningLimitSec = 1.0
)

// DBWrapper wrapper object to facilitate db interaction
type DBWrapper struct {
	Id string `bson:"_id"`
	types.Document
}

// HashifyDocument returns document with filled hash related parameters
func HashifyDocument(id string, doc types.Document,
	created, updated int64) (*types.Document, error) {
	doc.Document_id = id
	b, err := json.Marshal(doc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshall the document")
	}

	h := md5.New()
	if _, err := h.Write(b); err != nil {
		return nil, errors.Wrap(err, "unable to write to hash")
	}

	bs := h.Sum(nil)

	doc.Etag = fmt.Sprintf("%x", bs)
	doc.Version = types.DocumentVersion{
		Created: created,
		Updated: updated,
		Hash:    doc.Etag,
	}

	return &doc, nil
}
