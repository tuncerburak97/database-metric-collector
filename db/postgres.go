package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tuncerburak97/database-metric-collector/config"
	"github.com/tuncerburak97/database-metric-collector/metrics"
	"log"
	"time"
)

func MonitorPostgres(dbConfig config.DatabaseConfig, metrics *metrics.DatabaseMetrics, results chan<- bool) {
	dsn := config.BuildPostgresDSN(dbConfig.Config)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening PostgreSQL database %s: %v", dbConfig.Name, err)
	}
	log.Println("Connected to database: ", dbConfig.Name)
	defer db.Close()

	db.SetMaxIdleConns(dbConfig.Config.MaxIdleConns)
	db.SetMaxOpenConns(dbConfig.Config.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.Config.ConnMaxLifetime) * time.Minute)
	dbConfig.DSN = dsn
	MonitorDatabase(dbConfig, metrics, results)
}
