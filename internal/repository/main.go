package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Log    LogRepositoryInterface
}

func New() Repos {
	repos := Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{},
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
