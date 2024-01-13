package repository

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"entgo.io/ent/dialect"
	"github.com/enuesaa/pinit/pkg/ent"
	_ "github.com/mattn/go-sqlite3"
)

type DbRepositoryInterface interface {
	Open() error
	Close() error
	Appconf() *ent.AppconfClient
	Binder() *ent.BinderClient
	Note() *ent.NoteClient
	Action() *ent.ActionClient
	Migrate() error
}

type DbRepository struct {
	client *ent.Client
}

func (repo *DbRepository) Open() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dbPath := filepath.Join(homedir, ".pinit", "pinit.db")
	dsn := fmt.Sprintf("file:%s?_fk=1", dbPath)
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
