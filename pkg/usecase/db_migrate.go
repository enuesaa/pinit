package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func DBMigrate(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	is, err := binderSrv.IsTableExist()
	if err != nil {
		return fmt.Errorf("failed to migrate database because something error happens on pre check: %s", err.Error())
	}
	if is {
		return fmt.Errorf("failed to migrate database because binder table already exists.")
	}
	return repos.Db.Migrate()
}
