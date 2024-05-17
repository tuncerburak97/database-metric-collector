package metrics

import "github.com/prometheus/client_golang/prometheus"

type DatabaseMetrics struct {
	CPUUsage             prometheus.Gauge
	MemoryUsage          prometheus.Gauge
	TotalActiveQueries   *prometheus.GaugeVec
	AvgQueryDuration     *prometheus.GaugeVec
	QueryDurations       *prometheus.HistogramVec
	IndividualQueryTimes *prometheus.HistogramVec
	ConnectionErrors     *prometheus.CounterVec
	QueryErrors          *prometheus.CounterVec
	TransactionCount     *prometheus.CounterVec
	CacheUsage           *prometheus.GaugeVec
	DiskIO               *prometheus.GaugeVec
	NetworkTraffic       *prometheus.GaugeVec
	ConnectionPoolStats  *prometheus.GaugeVec
	DatabaseSize         *prometheus.GaugeVec
}

func NewDatabaseMetrics() *DatabaseMetrics {
	return &DatabaseMetrics{
		CPUUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "app_cpu_usage_percentage",
			Help: "CPU usage of the application.",
		}),
		MemoryUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "app_memory_usage_bytes",
			Help: "Memory usage of the application.",
		}),
		TotalActiveQueries: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "db_total_active_queries",
			Help: "Total number of active queries.",
		}, []string{"db_name"}),
		AvgQueryDuration: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "db_avg_query_duration_seconds",
			Help: "Average duration of database queries.",
		}, []string{"db_name"}),
		QueryDurations: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Duration of database queries.",
			Buckets: prometheus.DefBuckets,
		}, []string{"db_name"}),
		IndividualQueryTimes: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "db_individual_query_duration_seconds",
			Help:    "Execution times for individual queries.",
			Buckets: prometheus.DefBuckets,
		}, []string{"db_name", "query"}),
		ConnectionErrors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "db_connection_errors_total",
			Help: "Total number of database connection errors.",
		}, []string{"db_name"}),
		QueryErrors: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "db_query_errors_total",
			Help: "Total number of database query errors.",
		}, []string{"db_name"}),
		TransactionCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "db_transaction_count_total",
			Help: "Total number of database transactions.",
		}, []string{"db_name"}),
		CacheUsage: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "db_cache_usage_bytes",
			Help: "Cache usage of the database.",
		}, []string{"db_name"}),
		DiskIO: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "db_disk_io_bytes",
			Help: "Disk I/O of the database.",
		}, []string{"db_name"}),
		NetworkTraffic: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "db_network_traffic_bytes",
			Help: "Network traffic of the database.",
		}, []string{"db_name"}),
		ConnectionPoolStats: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "db_connection_pool_size",
			Help: "Connection pool size of the database.",
		}, []string{"db_name"}),
		DatabaseSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "db_size_bytes",
			Help: "Total size of the database.",
		}, []string{"db_name"}),
	}
}

func (dm *DatabaseMetrics) RecordQueryDuration(dbName string, duration float64) {
	dm.QueryDurations.WithLabelValues(dbName).Observe(duration)
}

func (dm *DatabaseMetrics) RecordIndividualQueryTime(dbName, query string, duration float64) {
	dm.IndividualQueryTimes.WithLabelValues(dbName, query).Observe(duration)
}

func (dm *DatabaseMetrics) SetCPUUsage(usage float64) {
	dm.CPUUsage.Set(usage)
}

func (dm *DatabaseMetrics) SetMemoryUsage(usage float64) {
	dm.MemoryUsage.Set(usage)
}

func (dm *DatabaseMetrics) SetTotalActiveQueries(dbName string, count float64) {
	dm.TotalActiveQueries.WithLabelValues(dbName).Set(count)
}

func (dm *DatabaseMetrics) SetAvgQueryDuration(dbName string, duration float64) {
	dm.AvgQueryDuration.WithLabelValues(dbName).Set(duration)
}

func (dm *DatabaseMetrics) IncrementConnectionErrors(dbName string) {
	dm.ConnectionErrors.WithLabelValues(dbName).Inc()
}

func (dm *DatabaseMetrics) IncrementQueryErrors(dbName string) {
	dm.QueryErrors.WithLabelValues(dbName).Inc()
}

func (dm *DatabaseMetrics) IncrementTransactionCount(dbName string) {
	dm.TransactionCount.WithLabelValues(dbName).Inc()
}

func (dm *DatabaseMetrics) SetCacheUsage(dbName string, usage float64) {
	dm.CacheUsage.WithLabelValues(dbName).Set(usage)
}

func (dm *DatabaseMetrics) SetDiskIO(dbName string, io float64) {
	dm.DiskIO.WithLabelValues(dbName).Set(io)
}

func (dm *DatabaseMetrics) SetNetworkTraffic(dbName string, traffic float64) {
	dm.NetworkTraffic.WithLabelValues(dbName).Set(traffic)
}

func (dm *DatabaseMetrics) SetConnectionPoolStats(dbName string, size float64) {
	dm.ConnectionPoolStats.WithLabelValues(dbName).Set(size)
}

func (dm *DatabaseMetrics) SetDatabaseSize(dbName string, size float64) {
	dm.DatabaseSize.WithLabelValues(dbName).Set(size)
}

func InitPrometheus() *DatabaseMetrics {
	dm := NewDatabaseMetrics()
	prometheus.MustRegister(dm.CPUUsage)
	prometheus.MustRegister(dm.MemoryUsage)
	prometheus.MustRegister(dm.TotalActiveQueries)
	prometheus.MustRegister(dm.AvgQueryDuration)
	prometheus.MustRegister(dm.QueryDurations)
	prometheus.MustRegister(dm.IndividualQueryTimes)
	prometheus.MustRegister(dm.ConnectionErrors)
	prometheus.MustRegister(dm.QueryErrors)
	prometheus.MustRegister(dm.TransactionCount)
	prometheus.MustRegister(dm.CacheUsage)
	prometheus.MustRegister(dm.DiskIO)
	prometheus.MustRegister(dm.NetworkTraffic)
	prometheus.MustRegister(dm.ConnectionPoolStats)
	prometheus.MustRegister(dm.DatabaseSize)
	return dm
}
