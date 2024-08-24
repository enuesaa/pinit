package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"modernc.org/sqlite"
)

func init() {
	driver := sqliteDriver{
		sqlite.Driver{},
	}
	sql.Register("sqlite3", &driver)
}

type sqliteDriver struct {
	sqlite.Driver
}

// see https://zenn.dev/nobonobo/articles/e9f17d183c19f6
func (d *sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	return conn, nil
}
