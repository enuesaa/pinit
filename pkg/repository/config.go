package repository

import (
	"github.com/spf13/viper"
)

//Deprecated
type ConfigRepositoryInterface interface {
	Init()
	Load() error
	DbHost() string
	SetDbHost(dbhost string)
	DbUsername() string
	SetDbUsername(dbusername string)
	DbPassword() string
	SetDbPassword(dbpassword string)
	DbName() string
	SetDbName(dbname string)
	ChatgptToken() string
	SetChatgptToken(chatgpttoken string)
	Write() error
}

type ConfigRepository struct{}

func (repo *ConfigRepository) Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.pinit")
}

func (repo *ConfigRepository) Load() error {
	return viper.ReadInConfig()
}

func (repo *ConfigRepository) DbHost() string {
	return viper.GetString("dbhost")
}

func (repo *ConfigRepository) SetDbHost(dbhost string) {
	viper.Set("dbhost", dbhost)
}

func (repo *ConfigRepository) DbUsername() string {
	return viper.GetString("dbusername")
}

func (repo *ConfigRepository) SetDbUsername(dbusername string) {
	viper.Set("dbusername", dbusername)
}

func (repo *ConfigRepository) DbPassword() string {
	return viper.GetString("dbpassword")
}

func (repo *ConfigRepository) SetDbPassword(dbpassword string) {
	viper.Set("dbpassword", dbpassword)
}

func (repo *ConfigRepository) DbName() string {
	return viper.GetString("dbname")
}

func (repo *ConfigRepository) SetDbName(dbname string) {
	viper.Set("dbname", dbname)
}

func (repo *ConfigRepository) ChatgptToken() string {
	return viper.GetString("chatgpttoken")
}

func (repo *ConfigRepository) SetChatgptToken(chatgpttoken string) {
	viper.Set("chatgpttoken", chatgpttoken)
}

func (repo *ConfigRepository) Write() error {
	return viper.WriteConfig()
}
