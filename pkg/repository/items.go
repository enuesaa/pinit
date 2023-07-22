package repository

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"golang.org/x/exp/slices"
)

type ItemInfo struct {
	Tag string
	Filename string
}

type Item struct {
	Tag string
	Filename string
	Content string 
}

type ItemsRepositoryInterface interface {
	CreateRegistry()
	GetRegistryPath() string
	GetItemPath(tag string, filename string) string
	ListItems() []ItemInfo
	FilterItemsByTags(items []ItemInfo, tags []string) []ItemInfo
	GetItem(tag string, filename string) (*Item, error)
	CreateItem(item Item)
	DeleteItem(tag string, filename string)
}

type ItemsRepository struct {}
func NewItemsRepository() *ItemsRepository {
	return &ItemsRepository{}
}

func (repo *ItemsRepository) CreateRegistry() {
	path := repo.GetRegistryPath()
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

func (repo *ItemsRepository) ListItems() []ItemInfo {
	path := repo.GetRegistryPath()
	entries, _ := os.ReadDir(path)

	r := regexp.MustCompile(`^([A-Za-z0-9]+)-(.*)$`)
	items := make([]ItemInfo, 0)
	for _, entry := range entries {
		// todo handle error
		matched := r.FindStringSubmatch(entry.Name())
		items = append(items, ItemInfo {
			Tag: matched[1],
			Filename: matched[2],
		})
	}
	return items
}

func (repo *ItemsRepository) FilterItemsByTags(items []ItemInfo, tags []string) []ItemInfo {
	list := make([]ItemInfo, 0)
	for _, item := range items {
		if slices.Contains(tags, item.Tag) {
			list = append(list, item)
		}
	}

	return list
}

func (repo *ItemsRepository) GetItem(tag string, filename string) (*Item, error) {
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
