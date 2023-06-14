package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
		logrus.Printf("Error reading config file: %v\n", err)
		return
	}

	// Create a new instance of the router
	app := fiber.New()

	// TODO: Use struct for any additional repo in future
	repos := initRepos(logger, cfg)

	// Initialize routes with the injected service
	srv := transport_tenants.Services{
		TenantService: tenants.NewTenantService(logger, cfg, repos),
	}

	transport_tenants.InitRoutes(app, srv)

	// Start the HTTP server
	port := cfg.MongoClient.Port
	logger.Printf("Server started on port %s\n", port)
	logger.Fatal(app.Listen(":" + port))
}

func initRepos(logger *logrus.Logger, cfg config.Config) repository.TenantRepository {
	mongoDB, err := mongo.NewMongoDB(cfg.MongoClient, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	return repository.NewTenantRepository(logger, cfg, mongoDB)
}

func configureLogger() *logrus.Logger {
	//Create a new Logrus logger
	logger := logrus.New()

	// Set the log level basedon the LOG_LEVEL environment variable
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	//Set the output to Stderr and enable colored output
	logger.SetOutput(os.Stderr)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	return logger
}
