package repository

import (
	"context"
	"fmt"

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
	config ConfigRepositoryInterface
	Tls    bool
}

func (repo *DbRepository) Open() error {
    client, err := ent.Open(dialect.SQLite, "file:ent.db?_fk=1")
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

func (repo *DbRepository) dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", repo.config.DbUsername(), repo.config.DbPassword(), repo.config.DbHost(), repo.config.DbName())
	params := "interpolateParams=true&parseTime=true"
	if repo.Tls {
		params = "interpolateParams=true&tls=true&parseTime=true"
	}
	return fmt.Sprintf("%s?%s", dsn, params)
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
