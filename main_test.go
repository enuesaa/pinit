package main

import (
	"log"
	"testing"
	// "github.com/enuesaa/pinit/pkg/repository"
)

func TestMain(m *testing.M) {
	// repos := repository.NewTestRepos()

	// if err := usecase.Migrate(repos); err != nil {
	// 	log.Fatalf("Error: failed to migrate, %s", err.Error())
	// }

	code := m.Run()
	log.Printf("test finished with code %d", code)
}
