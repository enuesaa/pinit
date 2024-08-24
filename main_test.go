package main

import (
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

func TestMain(m *testing.M) {
	repos := repository.New()
	dbPath := "test.db"
	repos.Db.Use(dbPath)

	defer func ()  {
		if err := repos.Fs.Remove(dbPath); err != nil {
			repos.Log.Errf(err, "failed to remove database file")
		}
	}()

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
