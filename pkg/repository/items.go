package repository

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type Item struct {
	Tag string
	Filename string
	Content string // todo 
}

type ItemsRepository struct {}

func (repo *ItemsRepository) CreateRegistry() error {
	path, err := GetRegistryPath()
	if err != nil {
		return err
	}
	os.Mkdir(path, 0755)
	return nil
}

func (repo *ItemsRepository) GetRegistryPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".pinit"), nil
}

func (repo *ItemsRepository) GetItemPath(tag string, filename string) string {
	registry, _ := repo.GetRegistryPath()
	return path.Join(registry, fmt.Sprintf("%s-%s", tag, filename))
}

func (repo *ItemsRepository) ListItems() {
	path, _ := repo.GetRegistryPath()
	entries, _ := os.ReadDir(path)
	fmt.Printf("%+v", entries)
}

func (repo *ItemsRepository) ListItemsByTags(tags []string) {
	path, _ := repo.GetRegistryPath()
	entries, _ := os.ReadDir(path)
	fmt.Printf("%+v", entries)
}

// should return item struct
func (repo *ItemsRepository) GetItem(tag string, filename string) (*Item, error){
	path := repo.GetItemPath(tag, filename)
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	item := Item {
		Tag: tag,
		Filename: filename,
		Content: string(file),
	}
	return &item, nil
}

func (repo *ItemsRepository) CreateItem(item Item) {
	// path := repo.GetItemPath(tag, filename)
}

func (repo *ItemsRepository) DeleteItem(tag string, filename string) {
	path := repo.GetItemPath(tag, filename)
	if _, err := os.Stat(path); err != nil {
		return;
	}
	os.Remove(path)
}
