package db

import (
	"database/sql"
	"github.com/tuncerburak97/database-metric-collector/config"
	"github.com/tuncerburak97/database-metric-collector/metrics"
	"log"
)

func MonitorOracle(dbConfig config.DatabaseConfig, metrics *metrics.DatabaseMetrics, results chan<- bool) {
	db, err := sql.Open("oci8", dbConfig.DSN)
	if err != nil {
		log.Fatalf("Error opening Oracle database %s: %v", dbConfig.Name, err)
	}
	defer db.Close()

	// Oracle specific monitoring logic here
	MonitorDatabase(dbConfig, metrics, results)
}
