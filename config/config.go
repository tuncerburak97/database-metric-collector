package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Databases []DatabaseConfig `mapstructure:"databases"`
}

type DatabaseConfig struct {
	Name   string         `mapstructure:"name"`
	Type   string         `mapstructure:"type"`
	DSN    string         `mapstructure:"dsn,omitempty"`
	Config PostgresConfig `mapstructure:"config,omitempty"`
}

type PostgresConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbName"`
	SSLMode         string `mapstructure:"sslMode"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

func ReadConfig(filename string) (*Config, error) {
	viper.SetConfigFile(filename)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func BuildPostgresDSN(cfg PostgresConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
}
