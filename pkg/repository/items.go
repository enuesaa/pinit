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
	Content string 
}

type ItemsRepository struct {}

func (repo *ItemsRepository) CreateRegistry() {
	path, _ := GetRegistryPath()
	os.Mkdir(path, 0755)
}

func (repo *ItemsRepository) GetRegistryPath() string {
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, ".pinit")
}

func (repo *ItemsRepository) GetItemPath(tag string, filename string) string {
	registry := repo.GetRegistryPath()
	return path.Join(registry, fmt.Sprintf("%s-%s", tag, filename))
}

func (repo *ItemsRepository) ListItems() {
	path := repo.GetRegistryPath()
	entries, _ := os.ReadDir(path)
	fmt.Printf("%+v", entries)
}

func (repo *ItemsRepository) ListItemsByTags(tags []string) {
	path := repo.GetRegistryPath()
	entries, _ := os.ReadDir(path)
	fmt.Printf("%+v", entries)
}

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
	path := repo.GetItemPath(item.Tag, item.Filename)

	f, _ := os.Create(path)
	defer f.Close()
	f.Write([]byte(item.Content))
}

func (repo *ItemsRepository) DeleteItem(tag string, filename string) {
	path := repo.GetItemPath(tag, filename)
	if _, err := os.Stat(path); err != nil {
		return;
	}
	os.Remove(path)
}
