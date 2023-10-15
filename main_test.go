package main

import (
	"log"
	"os"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
)

func TestMain(m *testing.M) {
	repos := repository.NewRepos()
	databaseDsn := os.Getenv("PINIT_TEST_DSN")
	if databaseDsn == "" {
		log.Fatalf("test execution error: environment variable PINIT_TEST_DSN is empty.")
	}
	repos.Database.WithDsn(databaseDsn)

	configSrv := service.NewConfigSevice(repos)
	if err := configSrv.Migration(); err != nil {
		log.Fatalf("test execution error: failed to migration")
	}
	configSrv.Migration()

    code := m.Run()
	log.Printf("text finished with code %d", code)
}