package main

import (
	"fmt"
	"log"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tuncerburak97/database-metric-collector/config"
	"github.com/tuncerburak97/database-metric-collector/db"
	"github.com/tuncerburak97/database-metric-collector/metrics"
)

func main() {
	// Read configuration
	config, err := config.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Initialize Prometheus metrics
	metrics := metrics.InitPrometheus()

	// Create a channel to receive the results
	results := make(chan bool)

	// Start monitoring each database in a separate goroutine
	for _, dbConfig := range config.Databases {
		switch dbConfig.Type {
		case "mysql":
			go db.MonitorMySQL(dbConfig, metrics, results)
		case "postgres":
			go db.MonitorPostgres(dbConfig, metrics, results)
		case "oracle":
			go db.MonitorOracle(dbConfig, metrics, results)
		default:
			log.Printf("Unsupported database type: %s", dbConfig.Type)
		}
	}

	// Fiber instance
	app := fiber.New()

	// Logger middleware
	app.Use(logger.New())
	prometheus := fiberprometheus.New("db-metrics-app")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Start server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(app.Listen(":8080"))
}
