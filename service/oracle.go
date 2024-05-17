package service

type OracleStatsService struct{}

func (o *OracleStatsService) GetActiveConnectionsQuery() string {
	return "SELECT count(*) FROM v$session"
}

func (o *OracleStatsService) GetCacheUsageQuery() string {
	return "" // Implement Oracle specific query
}

func (o *OracleStatsService) GetConnectionPoolSizeQuery() string {
	return "" // Implement Oracle specific query
}

func (o *OracleStatsService) GetDiskIOQuery() string {
	return "" // Implement Oracle specific query
}

func (o *OracleStatsService) GetNetworkTrafficQuery() string {
	return "" // Implement Oracle specific query
}

func (o *OracleStatsService) GetDatabaseSizeQuery() string {
	return "" // Implement Oracle specific query
}
