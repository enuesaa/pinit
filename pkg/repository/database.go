package repository

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/ent"
	"github.com/enuesaa/pinit/pkg/ent/hook"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseRepositoryInterface interface {
	Dsn() string
	EntDb() (*ent.Client, error)
	IsTableExist(name string) (bool, error)
	CreateTable(schema interface{}) error
	ListAll(data interface{}) error
	WhereFirst(data interface{}, query string, value interface{}) error
	Create(data interface{}) error
	Update(data interface{}) error
	Delete(data interface{}) error
	Count(data interface{}, query string, value string) (int64, error)
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

//Deprecated
func (repo *DatabaseRepository) db() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(repo.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func (repo *DatabaseRepository) EntDb() (*ent.Client, error) {
	db, err := ent.Open("mysql", repo.Dsn())
	if err != nil {
		return nil, err
	}
	// defer db.Close()
	db.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.BinderFunc(func(ctx context.Context, m *ent.BinderMutation) (ent.Value, error) {
			fmt.Printf("a\n")
			return next.Mutate(ctx, m)
		})
	}, ent.OpCreate))

	db.Intercept(
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				value, err := next.Query(ctx, query)
				fmt.Printf("selected: %+v\n", value)
				return value, err
			})
		}),
	)

	return db, nil
}

func (repo *DatabaseRepository) IsTableExist(name string) (bool, error) {
	db, err := repo.db()
	if err != nil {
		return false, err
	}
	var count int64
	if err := db.Table(name).Count(&count).Error; err != nil {
		return false, nil
	}
	return true, nil
}

func (repo *DatabaseRepository) CreateTable(schema interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	return db.Migrator().CreateTable(schema)
}

func (repo *DatabaseRepository) Create(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	return db.Create(data).Error
}

func (repo *DatabaseRepository) Update(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	return db.Updates(data).Error
}

func (repo *DatabaseRepository) Delete(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	return db.Delete(data).Error
}

func (repo *DatabaseRepository) ListAll(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	return db.Find(data).Error
}

func (repo *DatabaseRepository) WhereFirst(data interface{}, query string, value interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	return db.Where(query, value).First(data).Error
}

func (repo *DatabaseRepository) Count(data interface{}, query string, value string) (int64, error) {
	db, err := repo.db()
	if err != nil {
		return 0, err
	}
	var count int64
	if err := db.Model(data).Where(query, value).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
