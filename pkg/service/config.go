package service

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	DatabaseDsn string `toml:"database_dsn"`
}

type ConfigService struct {}

func (srv *ConfigService) SetupRegistry() error {
	return nil
}

func (srv *ConfigService) Read() Config {
	return Config{}
}

func (srv *ConfigService) Write() error {
	return nil
}


