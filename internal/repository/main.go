package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Log    LogRepositoryInterface
}

func New() (Repos, error) {
	repos := Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{},
		Log:    &LogRepository{},
	}
	return repos, nil
}

func NewMock() (Repos, error) {
	repos := Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{},
		Log:    &LogRepository{},
	}
	return repos, nil
}
