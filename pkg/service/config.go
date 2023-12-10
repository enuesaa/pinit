package service

import (
	"path/filepath"

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
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.pinit")

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
	homedir, err := srv.repos.Fs.HomeDir()
	if err != nil {
		return false
	}
	path := filepath.Join(homedir, ".pinit", "config.toml")

	return srv.repos.Fs.IsExist(path)
}

func (srv *ConfigService) Read() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (srv *ConfigService) Write(config Config) error {
	viper.Set("dbname", config.DbName)
	viper.Set("dbhost", config.DbHost)
	viper.Set("dbusername", config.DbUsername)
	viper.Set("dbpassword", config.DbPassword)
	viper.Set("chatgpttoken", config.ChatgptToken)
	return viper.WriteConfig()
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
