package repository

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseRepositoryInterface interface {}

type DatabaseRepository struct{
	Dsn string
}

func (repo *DatabaseRepository) db() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(repo.Dsn), &gorm.Config{})
}

// this method might drop table in the future.
func (repo *DatabaseRepository) Migrate(schema interface{}) error {
	db, err := repo.db()
	if err != nil {
		return err
	}
	// db.Migrator().CreateTable(&Note{})
	// db.Migrator().DropTable(&Note{})
	return db.AutoMigrate(schema)
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

