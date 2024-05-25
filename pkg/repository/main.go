package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

type Env struct {
	dbPath string
}

func New(dbPath string) Repos {
	env := Env{
		dbPath: dbPath,
	}
	return Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{env: env},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}
