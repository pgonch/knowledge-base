package util

import (
	"database/sql"
	"fmt"
	"time"

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

// TimeCounter counts time spent on db interaction
func TimeCounter(f func() string, warningLimit float64) {
	qStart := time.Now()
	query := f()
	duration := time.Since(qStart).Seconds()
	if duration > warningLimit {
		log.Warnf("Query: %s; took %f seconds", query,
			duration)
	}
}

// https://doxygen.postgresql.org/memutils_8h.html
const maxPSQLQueryLength = 0x3fffffff

// ArrayQuery returns query taking into account array size and operand
func ArrayQuery(queryPrefix, queryPostfixTemplate, arrayColumnName string,
	arraySize int, operand string, startInd, nPostFixParams int) (string, error) {
	sqlQuery := queryPrefix
	ind := startInd - 1
	if arraySize > 0 {
		sqlQuery += "AND ( "
		for i := 0; i < arraySize; i++ {
			ind++
			sqlQuery += fmt.Sprintf("$%d = ANY(%s) ", ind, arrayColumnName)
			if i != arraySize-1 {
				sqlQuery += fmt.Sprintf("%s ", operand)
			}
		}
		sqlQuery += " ) "
	}
	tmp := []interface{}{}
	for i := 0; i < nPostFixParams; i++ {
		ind++
		tmp = append(tmp, ind)
	}
	sqlQuery += fmt.Sprintf("ORDER BY created ASC LIMIT $%d OFFSET $%d",
		tmp...)

	if len(sqlQuery) > maxPSQLQueryLength {
		return "", fmt.Errorf("Too big query, specify smaller array length")
	}

	return sqlQuery, nil
}

// Close tries to close the rows and check the returned error
func Close(rows *sql.Rows) {
	if rows == nil {
		return
	}

	if err := rows.Close(); err != nil {
		log.Fatalf("Problem with closing rows: %v", err)
	}
}
