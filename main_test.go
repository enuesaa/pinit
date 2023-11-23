package main

import (
	"log"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

func TestMain(m *testing.M) {
	repos := repository.NewTestRepos()
	if repos.Database.GetDsn() == "" {
		log.Fatalf("test execution error: environment variable PINIT_TEST_DATABASE_DSN is empty.")
	}

	usecase.Migrate(repos)
	if err := usecase.Migrate(repos); err != nil {
		log.Fatalf("TestExecutionError: failed to migrate")
	}

	code := m.Run()
	log.Printf("text finished with code %d", code)
}
