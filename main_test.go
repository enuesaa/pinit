package main

import (
	"flag"
	"log"
	"testing"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
)

var databseDsnFlag = flag.String("dsn", "", "database dsn for test")

func TestMain(m *testing.M) {
	flag.Parse()
	if len(*databseDsnFlag) == 0 {
		log.Fatalf("Error: database dsn flag is required to run test.")
	}
	repos := repository.NewTestRepos(*databseDsnFlag)

	usecase.Migrate(repos)
	if err := usecase.Migrate(repos); err != nil {
		log.Fatalf("Error: failed to migrate")
	}

	code := m.Run()
	log.Printf("text finished with code %d", code)
}
