package repository

type Repos struct {
	Fs       FsRepositoryInterface
	Database DatabaseRepositoryInterface
	Prompt   PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Fs:       &FsRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}

func NewTestRepos() Repos {
	return Repos{
		Fs:       &FsMockRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}
