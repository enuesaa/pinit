package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func ListBinders(repos repository.Repos) []service.Binder {
	binderSrv := service.NewBinderService(repos)
	return binderSrv.List()
}

func CreateWithPrompt(repos repository.Repos) error {
	binderSrv := service.NewBinderService(repos)
	binder, err := binderSrv.RunPrompt(service.Binder{})
	if err != nil {
		return err
	}
	if err := binderSrv.Create(binder); err != nil {
		return err
	}

	noteSrv := service.NewNoteService(repos)
	note, err := noteSrv.RunPrompt(service.Note{ID: binder.ID})
	if err != nil {
		return err
	}
	return noteSrv.Create(&note)
}

func WriteNote(repos repository.Repos, binderName string) error {
	binderSrv := service.NewBinderService(repos)
	binder, err := binderSrv.GetByName(binderName)
	if err != nil {
		return err
	}

	noteSrv := service.NewNoteService(repos)
	note, err := noteSrv.RunPrompt(service.Note{ID: binder.ID})
	if err != nil {
		return err
	}
	return noteSrv.Create(&note)
}

func DescribeBinder(repos repository.Repos, binderName string) (service.Binder, error) {
	binderSrv := service.NewBinderService(repos)
	return binderSrv.GetByName(binderName)
}

func ListBinderNotes(repos repository.Repos, binderId string) ([]service.Note, error) {
	return make([]service.Note, 0), nil
}

//Deprecated
func DescribeFirstNote(repos repository.Repos, binderId uint) (service.Note, error) {
	noteSrv := service.NewNoteService(repos)
	return noteSrv.GetFirstByBinderId(binderId)
}

func Delete(repos repository.Repos, binderName string) error {
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
