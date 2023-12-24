package repository

type Repos struct {
	Config   ConfigRepositoryInterface
	Fs       FsRepositoryInterface
	Database DatabaseRepositoryInterface
	Prompt   PromptInterface
}

func NewRepos() Repos {
	config := ConfigRepository{}

	return Repos{
		Config: &config,
		Fs:     &FsRepository{},
		Database: &DatabaseRepository{
			config: &config,
			Tls: true,
		},
		Prompt: &Prompt{},
	}
}

func NewTestRepos() Repos {
	config := ConfigMockRepository{}

	return Repos{
		Config: &config,
		Fs:     &FsMockRepository{},
		Database: &DatabaseRepository{
			config: &config,
			Tls: false,
		},
		Prompt: &Prompt{},
	}
}
