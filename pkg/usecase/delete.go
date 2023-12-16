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

	// delete related notes.
	noteSrv := service.NewNoteService(repos)
	note, err := noteSrv.GetFirstByBinderId(binder.ID)
	if err != nil {
		return err
	}
	if err := noteSrv.Delete(note.ID); err != nil {
		return err
	}

	return binderSrv.Delete(binderName)
}
