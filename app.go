package main

import (
	"context"
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

type App struct {
	ctx context.Context
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet2(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ListBinders() ([]service.Binder, error) {
	binderSrv := service.NewBinderService(repository.NewRepos())
	return binderSrv.List()
}
