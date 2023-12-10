package repository

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbHost       string
	DbUsername   string
	DbPassword   string
	DbName       string
	ChatgptToken string
}

type ConfigRepositoryInterface interface {
	Init()
	Read() Config
	Write(config Config) error
}

type ConfigRepository struct {}

func (repo *ConfigRepository) Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.pinit")
}

func (repo *ConfigRepository) Read() Config {
	if err := viper.ReadInConfig(); err != nil {
		return Config{}
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}
	}
	return config
}

func (repo *ConfigRepository) Write(config Config) error {
	viper.Set("dbname", config.DbName)
	viper.Set("dbhost", config.DbHost)
	viper.Set("dbusername", config.DbUsername)
	viper.Set("dbpassword", config.DbPassword)
	viper.Set("chatgpttoken", config.ChatgptToken)
	return viper.WriteConfig()
}
