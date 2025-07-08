package app

import "go-product-app/common/postgresql"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	return &ConfigurationManager{PostgreSqlConfig: getPostgreSqlConfig()}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		UserName:              "postgres",
		Password:              "postgres",
		DatabaseName:          "productapp",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	}
}
