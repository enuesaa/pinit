package repository

import (
	"os"
	"path/filepath"
)

func CreateRegistry() error {
	path, err := GetRegistryPath()
	if err != nil {
		return err
	}
	os.Mkdir(path, 0755)
	return nil
}

func GetRegistryPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".pinit"), nil
}