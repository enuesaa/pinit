package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseRepositoryInterface interface {
	GetDsn() string
	WithTls(is bool)
	IsTableExist(name string) (bool, error)
	CreateTable(schema interface{}) error
	ListAll(data interface{}) error
	Where(data interface{}, results interface{}) error
	WhereFirst(data interface{}) error
	Create(data interface{}) error
	Update(data interface{}) error
	Delete(data interface{}) error
	Count(data interface{}, query string, value string) (int64, error)
}

type DatabaseRepository struct {
	config ConfigRepositoryInterface
	Tls    bool
}

func (repo *DatabaseRepository) GetDsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", repo.config.DbUsername(), repo.config.DbPassword(), repo.config.DbHost(), repo.config.DbName())
	params := "interpolateParams=true"
	if repo.Tls {
		params = "tls=true&interpolateParams=true"
	}

	return fmt.Sprintf("%s?%s", dsn, params)
}

func (repo *DatabaseRepository) WithTls(is bool) {
	repo.Tls = is
}

func (repo *DatabaseRepository) db() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(repo.GetDsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func (repo *DatabaseRepository) IsTableExist(name string) (bool, error) {
	db, err := repo.db()
	if err != nil {
		return false, err
	}
	var count int64
	if err := db.Table(name).Count(&count).Error; err != nil {
		return false, err
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

func (repo *DatabaseRepository) Where(data interface{}, results interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	var i int64
	db.Where(data).Count(&i)
	fmt.Printf("%+v%d", data, i)
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

func (repo *DatabaseRepository) Count(data interface{}, query string, value string) (int64, error) {
	db, err := repo.db()
	if err != nil {
		return 0, err
	}
	var count int64
	db.Model(&data).Where(query, value).Count(&count)
	return count, nil
}
