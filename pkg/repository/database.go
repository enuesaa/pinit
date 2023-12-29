package repository

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/ent"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseRepositoryInterface interface {
	Dsn() string
	EntDb() (*ent.Client, error)
	Migrate() error
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
