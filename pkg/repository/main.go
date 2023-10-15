package repository

import (
	"os"
)

type Repos struct {
	Fshome   FshomeRepositoryInterface
	Database DatabaseRepositoryInterface
	Prompt   PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Fshome:   &FshomeRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}

func NewTestRepos() Repos {
	databaseDsn := os.Getenv("PINIT_TEST_DSN")

	databaseRepo := &DatabaseRepository{}
	databaseRepo.WithDsn(databaseDsn)

	return Repos{
		Fshome:   &FshomeRepository{},
		Database: databaseRepo,
		Prompt:   &Prompt{},
	}
}
