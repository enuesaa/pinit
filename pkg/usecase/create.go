package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func CreateWithPrompt(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	var binder service.Binder
	if err := binderSrv.RunPrompt(&binder); err != nil {
		return err
	}
	if err := binderSrv.CheckNameAvailable(binder.Name); err != nil {
		return err
	}
	binderId, err := binderSrv.Create(binder)
	if err != nil {
		return err
	}

	noteSrv := service.NewNoteService(repos)
	note := service.Note{BinderId: binderId}
	if err := noteSrv.RunPrompt(&note); err != nil {
		return err
	}
	if _, err := noteSrv.Create(note); err != nil {
		return err
	}
	return nil
}

func WriteNewNoteWithPrompt(repos repository.Repos, binderName string) error {
	binderSrv := service.NewBinderService(repos)
	binder, err := binderSrv.GetByName(binderName)
	if err != nil {
		return err
	}

	noteSrv := service.NewNoteService(repos)
	latestNote, err := noteSrv.GetFirstByBinderId(binder.ID)
	if err != nil {
		return err
	}
	note := service.Note{BinderId: binder.ID, Content: latestNote.Content}

	if err := noteSrv.RunPrompt(&note); err != nil {
		return err
	}
	if _, err := noteSrv.Create(note); err != nil {
		return err
	}
	return nil
}
