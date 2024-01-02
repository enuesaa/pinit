package repository

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/ent"
	_ "github.com/go-sql-driver/mysql"
)

type DbRepositoryInterface interface {
	Open() error
	Close() error
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
	client, err := ent.Open("mysql", repo.dsn())
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
	db, err := ent.Open("mysql", repo.dsn())
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Schema.Create(context.Background())
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
