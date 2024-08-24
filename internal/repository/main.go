package repository

import (
	"os"
	"path/filepath"
)

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Log    LogRepositoryInterface
}

func New() (Repos, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return Repos{}, err
	}
	dbPath := filepath.Join(homedir, ".pinit", "pinit.db")
	repos := Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{dbPath: dbPath},
		Log:    &LogRepository{},
	}
	return repos, nil
}

func NewMock() (Repos, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return Repos{}, err
	}
	dbPath := filepath.Join(homedir, ".pinit", "pinitest.db")
	repos := Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{dbPath: dbPath},
		Log:    &LogRepository{},
	}
	return repos, nil
}
