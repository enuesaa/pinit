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

func New(dbPath string) Repos {
	env := Env{
		dbPath: os.Getenv("PINIT_DB_PATH"),
	}
	return Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{env: env},
		Prompt: &Prompt{},
	}
}
