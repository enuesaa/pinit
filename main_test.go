package main

import (
	"log"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

func TestMain(m *testing.M) {
	repos := repository.NewRepos()

	if err := usecase.DBSetup(repos); err != nil {
		log.Fatalf("Error: failed to migrate, %s", err.Error())
	}

	code := m.Run()
	log.Printf("test finished with code %d", code)
}
