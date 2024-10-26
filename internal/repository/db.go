package repository

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"entgo.io/ent/dialect"
)

type DbRepositoryInterface interface {
	IsDBExist() bool
	CreateDB() error
	Open() error
	Close() error
	Appconf() *ent.AppconfClient
	Binder() *ent.BinderClient
	Note() *ent.NoteClient
	Action() *ent.ActionClient
	Migrate() error
}

type DbRepository struct {
	dbPath string
	client *ent.Client
}

func (repo *DbRepository) IsDBExist() bool {
	if _, err := os.Stat(repo.dbPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *DbRepository) CreateDB() error {
	return os.MkdirAll(filepath.Dir(repo.dbPath), os.ModePerm)
}

func (repo *DbRepository) dsn() string {
	return fmt.Sprintf("file:%s?_fk=1", repo.dbPath)
}

func (repo *DbRepository) Open() error {
	dsn := repo.dsn()
	client, err := ent.Open(dialect.SQLite, dsn)
	if err != nil {
		return err
	}
	repo.client = client
	return nil
}

func (repo *DbRepository) Close() error {
	if repo.client == nil {
		return nil
	}
	return repo.client.Close()
}

func (repo *DbRepository) Migrate() error {
	if err := repo.Open(); err != nil {
		return err
	}
	return repo.client.Schema.Create(context.Background())
}

func (repo *DbRepository) Appconf() *ent.AppconfClient {
	return repo.client.Appconf
}

func (repo *DbRepository) Binder() *ent.BinderClient {
	return repo.client.Binder
}

func (repo *DbRepository) Note() *ent.NoteClient {
	return repo.client.Note
}

func (repo *DbRepository) Action() *ent.ActionClient {
	return repo.client.Action
}
