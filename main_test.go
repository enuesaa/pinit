package main

import (
	"log"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

func TestMain(m *testing.M) {
	dbPath := "test.db"
	repos := repository.New(dbPath)

	defer func ()  {
		if err := repos.Fs.Remove(dbPath); err != nil {
			log.Printf("failed to remove database file because `%s`", err.Error())
		}
	}()

	if err := usecase.DBSetup(repos); err != nil {
		log.Panicf("Error: failed to migrate database because `%s`", err.Error())
	}
	if err := usecase.DBOpen(repos); err != nil {
		log.Panicf("Error: failed to open database because `%s`", err.Error())
	}
	defer usecase.DBClose(repos)

	// run test
	code := m.Run()

	log.Printf("test finished with code %d", code)
}
