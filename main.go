package main

import (
	"github.com/enuesaa/pinit/internal/repository"
	"github.com/enuesaa/pinit/internal/usecase"
)

func main() {
	repos := repository.New()

	if err := usecase.Serve(repos, 8080); err != nil {
		repos.Log.Err(err)
	}
}
