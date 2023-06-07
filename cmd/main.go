package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/logutils"
	"github.com/toc/config"
	"github.com/toc/pkg/mongo-lib"
	"github.com/toc/pkg/tenants"
	repository "github.com/toc/repo/tenants"
	transport_tenants "github.com/toc/transport/tenants"
)

func main() {
	// Configure HashiCorp logger
	logger := configureLogger()

	// Specify the path to the YAML file
	filePath := "config/config.yaml"

	// Read the configuration from the YAML file
	cfg, err := config.ReadConfigFromFile(filePath)
	if err != nil {
		logger.Printf("Error reading config file: %v\n", err)
		return
	}

	// Create a new instance of the router
	router := mux.NewRouter()

	// TODO: Use struct for any additional repo in future
	repos := initRepos(logger, cfg)

	// Initialize routes with the injected service
	srv := transport_tenants.Services{
		TenantService: tenants.NewTenantService(logger, cfg, repos),
	}

	transport_tenants.InitRoutes(router, srv)

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	logger.Printf("Server started on port %s\n", port)
	logger.Fatal(http.ListenAndServe(":"+port, router))
}

func initRepos(logger *log.Logger, cfg config.Config) repository.TenantRepository {
	mongoDB, err := mongo.NewMongoDB(cfg.MongoClient)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return repository.NewTenantRepository(logger, cfg, mongoDB)
}

func configureLogger() *log.Logger {
	// Create a new log filter to configure the desired log level
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(os.Getenv("LOG_LEVEL")),
		Writer:   os.Stderr,
	}

	// Create a new logger and set the log filter
	logger := log.New(filter, "", log.LstdFlags)

	return logger
}
