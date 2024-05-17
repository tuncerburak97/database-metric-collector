package service

type MySQLStatsService struct{}

func (m *MySQLStatsService) GetActiveConnectionsQuery() string {
	return "SHOW STATUS WHERE `variable_name` = 'Threads_connected'"
}

func (m *MySQLStatsService) GetCacheUsageQuery() string {
	return "" // Implement MySQL specific query
}

func (m *MySQLStatsService) GetConnectionPoolSizeQuery() string {
	return "" // Implement MySQL specific query
}

func (m *MySQLStatsService) GetDiskIOQuery() string {
	return "" // Implement MySQL specific query
}

func (m *MySQLStatsService) GetNetworkTrafficQuery() string {
	return "" // Implement MySQL specific query
}

func (m *MySQLStatsService) GetDatabaseSizeQuery() string {
	return "" // Implement MySQL specific query
}
