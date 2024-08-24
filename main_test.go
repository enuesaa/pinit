package main

import (
	"log"
	"testing"

	"github.com/enuesaa/pinit/internal/repository"
	"github.com/enuesaa/pinit/internal/usecase"
)

func TestMain(m *testing.M) {
	repos, err := repository.NewMock()
	if err != nil {
		log.Panic("failed to start app")
	}

	if err := usecase.DBSetup(repos); err != nil {
		repos.Log.Errf(err, "failed to migrate database")
	}
	if err := usecase.DBOpen(repos); err != nil {
		repos.Log.Errf(err, "failed to open database")
	}
	defer usecase.DBClose(repos)

	// run test
	code := m.Run()

	repos.Log.Info("test finished with code %d", code)
}
