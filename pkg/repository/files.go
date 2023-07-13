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



type Item struct {
	Tags []string
	Filename string
	Content string // todo 
}

// filename 
// - .pinit/{tag}-{filename}
// example
// - .pinit/main-.gitignore
// - .pinit/dev-index.js

func ListItems() {}
func ListItemsByTags(tags []string) {}
func GetItem(tag string, filename string) {} // should return item struct
func CreateItem(item Item) {}
func DeleteItem(tag string, filename string) {}
