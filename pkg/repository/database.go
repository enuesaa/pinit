package repository

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/ent"
	"github.com/enuesaa/pinit/pkg/ent/predicate"
)

type DatabaseRepositoryInterface interface {
	Dsn() string
	EntDb() (*ent.Client, error)
	Migrate() error
	CountBinder(ps ...predicate.Binder) (int, error)
}

type DatabaseRepository struct {
	config ConfigRepositoryInterface
	Tls    bool
}

func (repo *DatabaseRepository) Dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", repo.config.DbUsername(), repo.config.DbPassword(), repo.config.DbHost(), repo.config.DbName())
	params := "interpolateParams=true&parseTime=true"
	if repo.Tls {
		params = "interpolateParams=true&tls=true&parseTime=true"
	}
	return fmt.Sprintf("%s?%s", dsn, params)
}

func (repo *DatabaseRepository) EntDb() (*ent.Client, error) {
	db, err := ent.Open("mysql", repo.Dsn())
	if err != nil {
		return nil, err
	}
	// defer db.Close()
	return db, nil
}

func (repo *DatabaseRepository) Migrate() error {
	db, err := repo.EntDb()
	if err != nil {
		return err
	}
	return db.Schema.Create(context.Background())
}

func (repo *DatabaseRepository) CountBinder(ps ...predicate.Binder) (int, error) {
	db, err := repo.EntDb()
	if err != nil {
		return 0, err
	}
	return db.Binder.Query().Where(ps...).Count(context.Background())
}

func (repo *DatabaseRepository) QueryBinderAll(ps ...predicate.Binder) error {
	db, err := repo.EntDb()
	if err != nil {
		return err
	}
	return db.Schema.Create(context.Background())
}

func (repo *DatabaseRepository) QueryBinderFirst(ps ...predicate.Binder) error {
	db, err := repo.EntDb()
	if err != nil {
		return err
	}
	return db.Schema.Create(context.Background())
}
