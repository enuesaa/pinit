package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

// setup database.
// - create database file
// - run migration
func DBSetup(repos repository.Repos) error {
	if err := repos.Db.CreateDB(); err != nil {
		return err
	}
	if err := DBOpen(repos); err != nil {
		return err
	}

	binderSrv := service.NewBinderService(repos)
	is, err := binderSrv.IsTableExist()
	if err != nil {
		return fmt.Errorf("failed to migrate database because something error happens on pre check: %s", err.Error())
	}
	if is {
		return fmt.Errorf("failed to migrate database because binder table already exists.")
	}
	if err := repos.Db.Migrate(); err != nil {
		return err
	}
	return nil
}
