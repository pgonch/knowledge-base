//go:generate go-raml server -l go --no-main --no-apidocs --dir . --ramlfile kb.raml --import-path github.com/pgonch/knowledge-base
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pgonch/knowledge-base/goraml"
	"github.com/pgonch/knowledge-base/util"

	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"

	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB

	dbUser            = os.Getenv("POSTGRES_USER")
	dbPassword        = os.Getenv("POSTGRES_PASSWORD")
	dbHost            = os.Getenv("POSTGRES_HOST")
	logLevel          = os.Getenv("LOG_LEVEL")
	port              = os.Getenv("PORT")
	maxRecordsInAPage = os.Getenv("MAX_RECORDS_IN_A_PAGE")

	dsn = fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", dbUser,
		dbPassword, dbHost)
)

const (
	defaultPort     = "5000"
	defaultLogLevel = "debug"
)

func init() {

	if port == "" {
		port = defaultPort
	}

	if logLevel == "" {
		logLevel = defaultLogLevel
	}

	if maxRecordsInAPage != "" {
		tmp, err := strconv.ParseInt(maxRecordsInAPage, 10, 64)
		if err != nil {
			fmt.Println("not a valid MAX_RECORDS_IN_A_PAGE value: ",
				maxRecordsInAPage)
			os.Exit(1)
		}
		util.LimitMax = int(tmp)
		util.LimitDefault = int(tmp)
	}

	// log setup
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	if err := util.ParseLogLevel(logLevel); err != nil {
		fmt.Println("not a valid log level: ", logLevel)
		os.Exit(1)
	}

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Probems with connecting to %s: %v", dsn, err)
	}

}

// LoggingMiddleware logs the time took for request handling
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		uri := r.RequestURI

		next.ServeHTTP(w, r)

		log.Debugf("uri: %s took %v", uri, time.Since(start))
	})
}

func main() {

	// input validator
	if err := validator.SetValidationFunc("multipleOf",
		goraml.MultipleOf); err != nil {
		log.Errorf("unable to set validation function: %v", err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	initRoutes(r)

	r.Use(LoggingMiddleware)

	log.Println("starting server on " + ":" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Errorf("unable to run the server: %v", err)
	}
}
