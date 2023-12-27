package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func CheckTableStatus(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	noteSrv := service.NewNoteService(repos)
	actionSrv := service.NewActionService(repos)

	isBinderTableExist, err := binderSrv.IsTableExist()
	if err != nil {
		return err
	}
	if !isBinderTableExist {
		return fmt.Errorf("Binders table does not exist")
	}

	isNoteTableExist, err := noteSrv.IsTableExist()
	if err != nil {
		return err
	}
	if !isNoteTableExist {
		return fmt.Errorf("Notes table does not exist")
	}

	isActionTableExist, err := actionSrv.IsTableExist()
	if err != nil {
		return err
	}
	if !isActionTableExist {
		return fmt.Errorf("Action table does not exist")
	}

	return nil
}

func Migrate(repos repository.Repos) error {
	return repos.Database.Migrate()
}
