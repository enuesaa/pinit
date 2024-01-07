package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func DescribeBinder(repos repository.Repos, binderName string) (service.Binder, error) {
	binderSrv := service.NewBinderService(repos)
	if err := binderSrv.CheckNameAvailable(binderName); err == nil {
		return service.Binder{}, fmt.Errorf("binder not found.")
	}
	return binderSrv.GetByName(binderName)
}

// TODO refactor
func DescribeLatestNote(repos repository.Repos, binderName string) (service.Note, error) {
	binderSrv := service.NewBinderService(repos)
	if err := binderSrv.CheckNameAvailable(binderName); err == nil {
		return service.Note{}, fmt.Errorf("binder not found.")
	}
	binder, err := binderSrv.GetByName(binderName)
	if err != nil {
		return service.Note{}, fmt.Errorf("binder not found.")
	}

	noteSrv := service.NewNoteService(repos)
	notes, err := noteSrv.ListByBinderId(binder.ID)
	if err != nil {
		return service.Note{}, err
	}
	if len(notes) > 0 {
		return notes[len(notes)-1], nil
	}

	return service.Note{}, fmt.Errorf("no notes found.")
}
