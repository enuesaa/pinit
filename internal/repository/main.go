package repository

import "os"

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Log    LogRepositoryInterface
}

func New() Repos {
	tableName := os.Getenv("PINIT_TABLE_NAME")
	if tableName == "" {
		tableName = "pinit"
	}

	repos := Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{
			TableName: tableName,
		},
		Log:    &LogRepository{},
	}
	return repos
}

func NewMock() Repos {
	repos := Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{},
		Log:    &LogRepository{},
	}
	return repos
}
