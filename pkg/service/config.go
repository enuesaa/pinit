package service

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/viper"
)

type Config struct {
	DbHost       string
	DbUsername   string
	DbPassword   string
	DbName       string
	ChatgptToken string
}

func NewConfigSevice(repos repository.Repos) ConfigService {
	return ConfigService{ repos: repos }
}

type ConfigService struct {
	repos repository.Repos
}

func (srv *ConfigService) Init() error {
	viper.SetConfigType("toml")

	config, err := srv.Read()
	if err != nil {
		return err
	}

	srv.repos.Database.WithDbHost(config.DbHost)
	srv.repos.Database.WithTls(true)
	srv.repos.Database.WithDbUsername(config.DbUsername)
	srv.repos.Database.WithDbPassword(config.DbPassword)
	srv.repos.Database.WithDbName(config.DbName)
	return nil
}

func (srv *ConfigService) IsConfigExist() bool {
	if _, err := srv.repos.Fshome.ReadFile(".pinit", "config.toml"); err != nil {
		return false
	}
	return true
}

func (srv *ConfigService) Read() (*Config, error) {
	content, err := srv.repos.Fshome.ReadFile(".pinit", "config.toml")
	if err != nil {
		return nil, err
	}
	viper.ReadConfig(bytes.NewReader(content))

	config := Config {
		DbHost: viper.GetString("db_host"),
		DbUsername: viper.GetString("db_username"),
		DbPassword: viper.GetString("db_password"),
		DbName: viper.GetString("db_name"),
		ChatgptToken: viper.GetString("chatgpt_token"),
	}
	return &config, nil
}

func (srv *ConfigService) Write(config Config) error {
	if !srv.repos.Fshome.IsRegistryExist(".pinit") {
		if err := srv.repos.Fshome.CreateRegistry(".pinit"); err != nil {
			return err
		}
	}
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(config); err != nil {
		return err
	}
	return srv.repos.Fshome.WriteFile(".pinit", "config.toml", buf.String())
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
