package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func ListActions(repos repository.Repos) ([]service.Action, error) {
	actionSrv := service.NewActionService(repos)
	return actionSrv.List()
}
