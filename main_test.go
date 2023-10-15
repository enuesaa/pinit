package main

import (
	"log"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func TestMain(m *testing.M) {
	repos := repository.NewTestRepos()
	if repos.Database.GetDsn() == "" {
		log.Fatalf("test execution error: environment variable PINIT_TEST_DATABASE_DSN is empty.")
	}

	configSrv := service.NewConfigSevice(repos)
	if err := configSrv.Migration(); err != nil {
		log.Fatalf("test execution error: failed to migration")
	}

	code := m.Run()
	log.Printf("text finished with code %d", code)
}
