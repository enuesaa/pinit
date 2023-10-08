package service

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/enuesaa/pinit/pkg/repository"
)

type Config struct {
	DatabaseDsn string `toml:"database_dsn"`
}

type ConfigService struct {
	repos          repository.Repos
	registryName   string
	configFilename string
}

func NewConfigSevice(repos repository.Repos) ConfigService {
	return ConfigService{
		repos:          repos,
		registryName:   ".pinit",
		configFilename: "config.toml",
	}
}

func (srv *ConfigService) Read() (*Config, error) {
	content, err := srv.repos.Fshome.ReadFile(srv.registryName, srv.configFilename)
	if err != nil {
		return nil, err
	}
	var config Config
	if _, err := toml.Decode(content, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (srv *ConfigService) Write(config Config) error {
	if !srv.repos.Fshome.IsRegistryExist(srv.registryName) {
		if err := srv.repos.Fshome.CreateRegistry(srv.registryName); err != nil {
			return err
		}
	}
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(config); err != nil {
		return err
	}
	return srv.repos.Fshome.WriteFile(srv.registryName, srv.configFilename, buf.String())
}

func (srv *ConfigService) Init() error {
	config, err := srv.Read()
	if err != nil {
		return err
	}
	srv.repos.Database.WithDsn(config.DatabaseDsn)
	return nil
}
