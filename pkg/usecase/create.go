package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func CreateWithPrompt(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	binder, err := binderSrv.RunPrompt(service.Binder{})
	if err != nil {
		return err
	}
	if err := binderSrv.CheckNameAvailable(binder.Name); err != nil {
		return err
	}
	if err := binderSrv.Create(binder); err != nil {
		return err
	}

	noteSrv := service.NewNoteService(repos)
	note, err := noteSrv.RunPrompt(service.Note{BinderId: binder.ID})
	if err != nil {
		return err
	}
	return noteSrv.Create(&note)
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
	note, err := noteSrv.RunPrompt(service.Note{BinderId: binder.ID, Content: latestNote.Content})
	if err != nil {
		return err
	}
	return noteSrv.Create(&note)
}
