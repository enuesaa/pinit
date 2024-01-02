package repository

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/ent"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseRepositoryInterface interface {
	Open() error
	Close() error
	Binder() *ent.BinderClient
	Note() *ent.NoteClient
	Action() *ent.ActionClient
	Db() (*ent.Client, error)
	Migrate() error
}

type DatabaseRepository struct {
	client *ent.Client
	config ConfigRepositoryInterface
	Tls    bool
}

func (repo *DatabaseRepository) Open() error {
	client, err := ent.Open("mysql", repo.dsn())
	if err != nil {
		return err
	}
	repo.client = client
	return nil
}

func (repo *DatabaseRepository) Close() error {
	if repo.client == nil {
		return nil
	}
	return repo.client.Close()
}

func (repo *DatabaseRepository) dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", repo.config.DbUsername(), repo.config.DbPassword(), repo.config.DbHost(), repo.config.DbName())
	params := "interpolateParams=true&parseTime=true"
	if repo.Tls {
		params = "interpolateParams=true&tls=true&parseTime=true"
	}
	return fmt.Sprintf("%s?%s", dsn, params)
}

//Deprecated
func (repo *DatabaseRepository) Db() (*ent.Client, error) {
	db, err := ent.Open("mysql", repo.dsn())
	if err != nil {
		return nil, err
	}
	// defer db.Close()
	return db, nil
}

func (repo *DatabaseRepository) Migrate() error {
	db, err := repo.Db()
	if err != nil {
		return err
	}
	return db.Schema.Create(context.Background())
}

func (repo *DatabaseRepository) Binder() *ent.BinderClient {
	return repo.client.Binder
}

func (repo *DatabaseRepository) Note() *ent.NoteClient {
	return repo.client.Note
}

func (repo *DatabaseRepository) Action() *ent.ActionClient {
	return repo.client.Action
}
