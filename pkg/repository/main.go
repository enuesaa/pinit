package repository

type Repos struct {
	Config ConfigRepositoryInterface
	Fs     FsRepositoryInterface
	Db     DbRepositoryInterface
	Prompt PromptInterface
}

func NewRepos() Repos {
	config := ConfigRepository{}

	return Repos{
		Config: &config,
		Fs:     &FsRepository{},
		Db: &DbRepository{
			config: &config,
			Tls:    true,
		},
		Prompt: &Prompt{},
	}
}

func NewTestRepos() Repos {
	config := ConfigMockRepository{}

	return Repos{
		Config: &config,
		Fs:     &FsMockRepository{},
		Db: &DbRepository{
			config: &config,
			Tls:    false,
		},
		Prompt: &Prompt{},
	}
}
