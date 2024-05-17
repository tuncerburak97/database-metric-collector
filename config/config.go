package config

import (
	"io/ioutil"
)

type Config struct {
	Databases []DatabaseConfig `yaml:"databases"`
}

type DatabaseConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"`
}

func ReadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
