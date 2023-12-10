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

func (repo *Config) Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.pinit")
}

func (repo *Config) Read() {
	if err := viper.ReadInConfig(); err != nil {
		return
	}
	if err := viper.Unmarshal(repo); err != nil {
		return
	}
}

func (repo *Config) Write() error {
	viper.Set("dbname", repo.DbName)
	viper.Set("dbhost", repo.DbHost)
	viper.Set("dbusername", repo.DbUsername)
	viper.Set("dbpassword", repo.DbPassword)
	viper.Set("chatgpttoken", repo.ChatgptToken)
	return viper.WriteConfig()
}
