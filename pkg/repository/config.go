package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type FshomeRepositoryInterface interface {
	IsRegistryExist(registryName string) bool
	CreateRegistry(registryName string) error
	IsFileExist(registryName string, path string) bool
	WriteFile(registryName string, path string, content string) error
	ReadFile(registryName string, path string) (string, error)
}
type FshomeRepository struct {}

func (repo *FshomeRepository) isFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FshomeRepository) homedir() (string, error) {
	return os.UserHomeDir()
}

func (repo *FshomeRepository) IsRegistryExist(registryName string) bool {
	homedir, err := repo.homedir()
	if err != nil {
		return false
	}
	path := filepath.Join(homedir, registryName)
	return repo.isFileOrDirExist(path)
}

func (repo *FshomeRepository) CreateRegistry(registryName string) error {
	homedir, err := repo.homedir()
	if err != nil {
		return err
	}
	path := filepath.Join(homedir, registryName)
	return os.Mkdir(path, 0755)
}

func (repo *FshomeRepository) IsFileExist(registryName string, path string) bool {
	if !repo.IsRegistryExist(registryName) {
		return false
	}
	homedir, err := repo.homedir()
	if err != nil {
		return false
	}
	fullpath := filepath.Join(homedir, registryName, path)
	return repo.isFileOrDirExist(fullpath)
}

func (repo *FshomeRepository) WriteFile(registryName string, path string, content string) error {
	if !repo.IsRegistryExist(registryName) {
		return fmt.Errorf("registry does not exist.")
	}
	homedir, err := repo.homedir()
	if err != nil {
		return err
	}
	fullpath := filepath.Join(homedir, registryName, path)
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write([]byte(content)); err != nil {
		return err
	}
	return nil
}

func (repo *FshomeRepository) ReadFile(registryName string, path string) (string, error) {
	if !repo.IsFileExist(registryName, path) {
		return "", fmt.Errorf("file does not exist.")
	}
	homedir, err := repo.homedir()
	if err != nil {
		return "", err
	}
	fullpath := filepath.Join(homedir, registryName, path)
	content, err := os.ReadFile(fullpath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}



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
