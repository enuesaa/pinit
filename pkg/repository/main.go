package repository

type Repos struct {
	Config   *Config
	Fs       FsRepositoryInterface
	Database DatabaseRepositoryInterface
	Prompt   PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Config:   &Config{},
		Fs:       &FsRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}

func NewTestRepos() Repos {
	return Repos{
		Config:   &Config{},
		Fs:       &FsMockRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}
