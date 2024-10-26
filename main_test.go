package main

import (
	"testing"

	"github.com/enuesaa/pinit/internal/repository"
)

func TestMain(m *testing.M) {
	repos := repository.NewMock()

	// run test
	code := m.Run()

	repos.Log.Info("test finished with code %d", code)
}
