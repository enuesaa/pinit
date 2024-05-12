package main

import (
	"context"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

type App struct {
	ctx context.Context
	repos repository.Repos
}

func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
	if err := usecase.OpenDb(a.repos); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
func (a *App) shutdown(ctx context.Context) {
	usecase.CloseDb(a.repos)
}

func (a *App) ctl() *usecase.ServeCtl {
	ctl := usecase.NewServeCtl(a.repos)
	return &ctl
}

func (a *App) ListBinders() ([]usecase.ListBindersItem, error) {
	return a.ctl().ListBinders()
}
func (a *App) CreateBinder(req usecase.CreateBinderRequest) (usecase.ServeCreateResponse, error) {
	return a.ctl().CreateBinder(req)
}
func (a *App) DeleteBinder(binderId int) error {
	return a.ctl().DeleteBinder(binderId)
}

func (a *App) ListNotes(binderId int) ([]usecase.ListNotesItem, error) {
	return a.ctl().ListNotes(binderId)
}
func (a *App) CreateNote(req usecase.CreateNoteRequest) (usecase.ServeCreateResponse, error) {
	return a.ctl().CreateNote(req)
}

func (a *App) ListActios() ([]usecase.ListActionsItem, error) {
	return a.ctl().ListActions()
}

func (a *App) Chat(req usecase.ChatRequest) (usecase.ChatResponse, error) {
	return a.ctl().Chat(req)
}

func (a *App) Recog(req string) (usecase.RecogResponse, error) {
	return a.ctl().Recog([]byte(req))
}
