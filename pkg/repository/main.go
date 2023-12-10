package repository

type Repos struct {
	Config   ConfigRepositoryInterface
	Fs       FsRepositoryInterface
	Database DatabaseRepositoryInterface
	Prompt   PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Config:   &ConfigRepository{},
		Fs:       &FsRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}

func NewTestRepos() Repos {
	return Repos{
		Config:   &ConfigRepository{},
		Fs:       &FsMockRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}
