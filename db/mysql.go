package db

import (
	"database/sql"
	"github.com/tuncerburak97/database-metric-collector/config"
	"github.com/tuncerburak97/database-metric-collector/metrics"
	"log"
)

func MonitorMySQL(dbConfig config.DatabaseConfig, metrics *metrics.DatabaseMetrics, results chan<- bool) {
	db, err := sql.Open("mysql", dbConfig.DSN)
	if err != nil {
		log.Fatalf("Error opening MySQL database %s: %v", dbConfig.Name, err)
	}
	defer db.Close()

	// MySQL specific monitoring logic here
	MonitorDatabase(dbConfig, metrics, results)
}
