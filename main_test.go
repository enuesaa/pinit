package main

import (
	"log"
	"testing"

	"github.com/enuesaa/pinit/internal/repository"
)

func TestMain(m *testing.M) {
	repos, err := repository.NewMock()
	if err != nil {
		log.Panic("failed to start app")
	}

	// run test
	code := m.Run()

	repos.Log.Info("test finished with code %d", code)
}
