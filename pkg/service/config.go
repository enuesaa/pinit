package service

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/enuesaa/pinit/pkg/repository"
)

type Config struct {
	DbHost string `toml:"db_host"`
	DbUsername string `toml:"db_username"`
	DbPassword string `toml:"db_password"`
	DbName string `toml:"db_name"`
	ChatgptToken string `toml:"chatgpt_token"`
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

func (srv *ConfigService) IsConfigExist() bool {
	if _, err := srv.repos.Fshome.ReadFile(srv.registryName, srv.configFilename); err != nil {
		return false
	}
	return true
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

func (srv *ConfigService) RunPrompt(config Config) (*Config, error) {
	dbHost, err := srv.repos.Prompt.Ask("DB Host", config.DbHost)
	if err != nil {
		return nil, err
	}
	config.DbHost = dbHost
	dbUsername, err := srv.repos.Prompt.Ask("DB Username", config.DbUsername)
	if err != nil {
		return nil, err
	}
	config.DbUsername = dbUsername
	dbPassword, err := srv.repos.Prompt.Ask("DB Password", config.DbPassword)
	if err != nil {
		return nil, err
	}
	config.DbPassword = dbPassword
	dbName, err := srv.repos.Prompt.Ask("DB Name", config.DbName)
	if err != nil {
		return nil, err
	}
	config.DbName = dbName

	chatgptToken, err := srv.repos.Prompt.Ask("OpenAI API Token", config.ChatgptToken)
	if err != nil {
		return nil, err
	}
	config.ChatgptToken = chatgptToken

	return &config, nil
}

func (srv *ConfigService) GetDatabaseDsn() (string, error) {
	config, err := srv.Read()
	if err != nil {
		return "", err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbName)
	params := "tls=true&interpolateParams=true"

	return fmt.Sprintf("%s?%s", dsn, params), nil
}

func (srv *ConfigService) Init() error {
	dsn, err := srv.GetDatabaseDsn()
	if err != nil {
		return err
	}
	srv.repos.Database.WithDsn(dsn)
	return nil
}
