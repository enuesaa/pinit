package repository

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseRepositoryInterface interface {
	WithDsn(dsn string)
	IsTableExist(schema interface{}) (bool, error)
	CreateTable(schema interface{}) error
	ListAll(data interface{}) error
	WhereFirst(data interface{}) error
	Create(data interface{}) error
	Update(data interface{}) error
	Delete(data interface{}) error
}

type DatabaseRepository struct {
	Dsn string
}

func (repo *DatabaseRepository) WithDsn(dsn string) {
	repo.Dsn = dsn
}

func (repo *DatabaseRepository) db() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(repo.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func (repo *DatabaseRepository) IsTableExist(schema interface{}) (bool, error) {
	db, err := repo.db()
	if err != nil {
		return false, err
	}
	var count int64
	result := db.Model(schema).Count(&count)
	countErr := result.Error

	return countErr == nil, nil
}

func (repo *DatabaseRepository) CreateTable(schema interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	return db.Migrator().CreateTable(schema)
}

// example `show tables;`
func (repo *DatabaseRepository) RunRawSql(sql string) (*sql.Rows, error) {
	db, err := repo.db()
	if err != nil {
		return nil, err
	}
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rows, nil
}

func (repo *DatabaseRepository) Create(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	result := db.Create(data)
	return result.Error
}

func (repo *DatabaseRepository) Update(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	result := db.Updates(data)
	return result.Error
}

func (repo *DatabaseRepository) Delete(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	result := db.Delete(data)
	return result.Error
}

func (repo *DatabaseRepository) ListAll(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	db.Find(data)
	return nil
}

func (repo *DatabaseRepository) WhereFirst(data interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	db.Where(data).First(&data)
	return nil
}
