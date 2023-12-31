package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func DeleteBinder(repos repository.Repos, binderName string) error {
	binderSrv := service.NewBinderService(repos)
	binder, err := binderSrv.GetByName(binderName)
	if err != nil {
		return err
	}

	noteSrv := service.NewNoteService(repos)
	if err := noteSrv.DeleteByBinderId(binder.ID); err != nil {
		return err
	}

	return binderSrv.Delete(binder.ID)
}
