package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}
