package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func ListBinders(repos repository.Repos) ([]service.Binder, error) {
	binderSrv := service.NewBinderService(repos)
	return binderSrv.List()
}

func ListBinderNotes(repos repository.Repos, binderId uint) ([]service.Note, error) {
	noteSrv := service.NewNoteService(repos)
	return noteSrv.ListByBinderId(binderId)
}

func ListActions(repos repository.Repos) ([]service.Action, error) {
	actionSrv := service.NewActionService(repos)
	return actionSrv.List()
}
