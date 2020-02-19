//go:generate go-raml server -l go --no-main --no-apidocs --dir . --ramlfile kb.raml --import-path github.com/pgonch/knowledge-base
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/pgonch/knowledge-base/goraml"
	"github.com/pgonch/knowledge-base/handlers/document"
	"github.com/pgonch/knowledge-base/handlers/documents"
	"github.com/pgonch/knowledge-base/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"

	log "github.com/sirupsen/logrus"
)

var (
	dbHost   = os.Getenv("MONGO_HOST")
	logLevel = os.Getenv("LOG_LEVEL")
	port     = os.Getenv("PORT")
)

const (
	defaultPort     = "5000"
	defaultLogLevel = "debug"
	defaultDBHost   = "localhost"
)

func init() {

	if port == "" {
		port = defaultPort
	}

	if logLevel == "" {
		logLevel = defaultLogLevel
	}

	if dbHost == "" {
		dbHost = defaultDBHost
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

	dbClient, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017",
		dbHost)))
	if err != nil {
		log.Fatalf("unable to create client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer func(ctx context.Context) {
		if err := dbClient.Disconnect(ctx); err != nil {
			log.Fatalf("unable to disconnect: %v", err)
		}
	}(ctx)

	// could be as separate vars
	db := dbClient.Database("knowledgebase")
	dcts := db.Collection("documents")

	// input validator
	if err := validator.SetValidationFunc("multipleOf",
		goraml.MultipleOf); err != nil {
		log.Errorf("unable to set validation function: %v", err)
		os.Exit(1)
	}

	documents.SetDB(dcts)
	document.SetDB(dcts)
	document.SetDBClient(dbClient)

	r := mux.NewRouter()

	initRoutes(r)

	r.Use(LoggingMiddleware)

	log.Println("starting server on " + ":" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Errorf("unable to run the server: %v", err)
	}
}
