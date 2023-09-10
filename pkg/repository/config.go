package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DatabaseDsn string `toml:"database_dsn"`
}

type ConfigRepositoryInterface interface {
	GetRegistryPath() (string, error)
	CreateRegistry() error
	GetConfigFilePath() (string, error)
	HasConfigFile() bool
	WriteConfig(config Config) error
	ReadConfig() (*Config, error)
}
type ConfigRepository struct{}

func (repo *ConfigRepository) isFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *ConfigRepository) GetRegistryPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".pinit"), nil
}

func (repo *ConfigRepository) CreateRegistry() error {
	path, err := repo.GetRegistryPath()
	if err != nil {
		return err
	}
	if repo.isFileOrDirExist(path) {
		return nil
	}
	return os.Mkdir(path, 0755)
}

func (repo *ConfigRepository) GetConfigFilePath() (string, error) {
	registryPath, err := repo.GetRegistryPath()
	if err != nil {
		return "", err
	}
	path := filepath.Join(registryPath, "config.toml")
	return path, nil
}

func (repo *ConfigRepository) HasConfigFile() bool {
	path, err := repo.GetConfigFilePath()
	if err != nil {
		return false
	}
	return repo.isFileOrDirExist(path)
}

func (repo *ConfigRepository) WriteConfig(config Config) error {
	if err := repo.CreateRegistry(); err != nil {
		return err
	}

	path, err := repo.GetConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if err != nil {
		return err
	}

	return toml.NewEncoder(file).Encode(config)
}

func (repo *ConfigRepository) ReadConfig() (*Config, error) {
	if !repo.HasConfigFile() {
		return nil, fmt.Errorf("failed to read config file")
	}

	path, err := repo.GetConfigFilePath()
	if err != nil {
		return nil, err
	}
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
