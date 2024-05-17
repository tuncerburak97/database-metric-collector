package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/tuncerburak97/database-metric-collector/config"
	"github.com/tuncerburak97/database-metric-collector/metrics"
	"github.com/tuncerburak97/database-metric-collector/service"
	"log"
	"runtime"
	"time"
)

type DatabaseStatsInterface interface {
	GetActiveConnectionsQuery() string
	GetCacheUsageQuery() string
	GetConnectionPoolSizeQuery() string
	GetDiskIOQuery() string
	GetNetworkTrafficQuery() string
	GetDatabaseSizeQuery() string
}

func MonitorDatabase(dbConfig config.DatabaseConfig, metrics *metrics.DatabaseMetrics, results chan<- bool) {
	var dbStats DatabaseStatsInterface
	switch dbConfig.Type {
	case "mysql":
		dbStats = &service.MySQLStatsService{}
	case "postgres":
		dbStats = &service.PostgresStatsService{}
	case "oracle":
		dbStats = &service.OracleStatsService{}
	default:
		log.Fatalf("Unsupported database type: %s", dbConfig.Type)
	}

	dsn := config.BuildPostgresDSN(dbConfig.Config) // Replace this with appropriate DSN builder
	db, err := sql.Open(dbConfig.Type, dsn)
	if err != nil {
		log.Fatalf("Error opening database %s: %v", dbConfig.Name, err)
	}
	defer db.Close()

	db.SetMaxIdleConns(dbConfig.Config.MaxIdleConns)
	db.SetMaxOpenConns(dbConfig.Config.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(dbConfig.Config.ConnMaxLifetime) * time.Minute)

	for {
		start := time.Now()
		query := "SELECT 1"
		if _, err := db.Exec(query); err != nil {
			log.Printf("Error running query on database %s: %v", dbConfig.Name, err)
			metrics.IncrementQueryErrors(dbConfig.Name)
		}
		duration := time.Since(start).Seconds()

		metrics.RecordQueryDuration(dbConfig.Name, duration)
		metrics.RecordIndividualQueryTime(dbConfig.Name, query, duration)

		activeConnectionsQuery := dbStats.GetActiveConnectionsQuery()
		var activeConnections int
		err = db.QueryRow(activeConnectionsQuery).Scan(&activeConnections)
		if err != nil {
			log.Printf("Error getting active connections for database %s: %v", dbConfig.Name, err)
			metrics.IncrementConnectionErrors(dbConfig.Name)
		}
		metrics.SetTotalActiveQueries(dbConfig.Name, float64(activeConnections))

		// Update CPU and Memory usage
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		metrics.SetMemoryUsage(float64(m.Alloc))

		// Simulating CPU usage collection as an example
		cpuUsage := 10.0 // Placeholder value, replace with actual CPU usage collection method
		metrics.SetCPUUsage(cpuUsage)

		cacheUsageQuery := dbStats.GetCacheUsageQuery()
		var cacheUsage float64
		err = db.QueryRow(cacheUsageQuery).Scan(&cacheUsage)
		if err != nil {
			log.Printf("Error getting cache usage for database %s: %v", dbConfig.Name, err)
			metrics.IncrementConnectionErrors(dbConfig.Name)
		}
		metrics.SetCacheUsage(dbConfig.Name, cacheUsage) // 500 MB cache usage

		discIOQuery := dbStats.GetDiskIOQuery()
		var diskIO float64
		err = db.QueryRow(discIOQuery).Scan(&diskIO)
		if err != nil {
			log.Printf("Error getting disk I/O for database %s: %v", dbConfig.Name, err)
			metrics.IncrementConnectionErrors(dbConfig.Name)
		}
		metrics.SetDiskIO(dbConfig.Name, diskIO) // 100 MB/s disk I/O

		networkTrafficQuery := dbStats.GetNetworkTrafficQuery()
		var networkTraffic float64
		err = db.QueryRow(networkTrafficQuery).Scan(&networkTraffic)
		if err != nil {
			log.Printf("Error getting network traffic for database %s: %v", dbConfig.Name, err)
			metrics.IncrementConnectionErrors(dbConfig.Name)
		}
		metrics.SetNetworkTraffic(dbConfig.Name, networkTraffic) // 200 MB/s network traffic

		connectionPoolSizeQuery := dbStats.GetConnectionPoolSizeQuery()
		var connectionPoolSize float64
		err = db.QueryRow(connectionPoolSizeQuery).Scan(&connectionPoolSize)
		if err != nil {
			log.Printf("Error getting connection pool size for database %s: %v", dbConfig.Name, err)
			metrics.IncrementConnectionErrors(dbConfig.Name)
		}
		metrics.SetConnectionPoolStats(dbConfig.Name, connectionPoolSize)

		databaseSizeQuery := dbStats.GetDatabaseSizeQuery()
		var databaseSize float64
		err = db.QueryRow(databaseSizeQuery).Scan(&databaseSize)
		if err != nil {
			log.Printf("Error getting database size for database %s: %v", dbConfig.Name, err)
			metrics.IncrementConnectionErrors(dbConfig.Name)
		}
		metrics.SetDatabaseSize(dbConfig.Name, databaseSize)

		metrics.IncrementTransactionCount(dbConfig.Name)

		time.Sleep(10 * time.Second)
		results <- true
	}
}
