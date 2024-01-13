package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Prompt PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Fs:     &FsRepository{},
		Db:     &DbRepository{},
		Prompt: &Prompt{},
	}
}

func NewTestRepos() Repos {
	return Repos{
		Fs:     &FsMockRepository{},
		Db:     &DbRepository{},
		Prompt: &Prompt{},
	}
}
