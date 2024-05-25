package repository

import (
	"os"
)

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Prompt PromptInterface
}

type Env struct {
	dbPath string
}

func NewRepos() Repos {
	env := Env{
		dbPath: os.Getenv("PINIT_DB_PATH"),
	}
	return Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{env: env},
		Prompt: &Prompt{},
	}
}

func NewTestRepos() Repos {
	return Repos{
		Fs:     &FsMockRepository{},
		Db:     &DbRepository{},
		Prompt: &Prompt{},
	}
}
