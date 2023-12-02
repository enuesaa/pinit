package main

import (
	"flag"
	"log"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

var testDbHostFlag = flag.String("dbhost", "", "database host for test")

func TestMain(m *testing.M) {
	flag.Parse()
	if len(*testDbHostFlag) == 0 {
		log.Fatalf("Error: database host flag is required to run test.")
	}
	repos := repository.NewTestRepos(*testDbHostFlag)

	usecase.Migrate(repos)
	if err := usecase.Migrate(repos); err != nil {
		log.Fatalf("Error: failed to migrate")
	}

	code := m.Run()
	log.Printf("text finished with code %d", code)
}
