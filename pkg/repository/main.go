package repository

type Repos struct {
	Fs       FsRepositoryInterface
	Fshome   FshomeRepositoryInterface
	Database DatabaseRepositoryInterface
	Prompt   PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Fs:       &FsRepository{},
		Fshome:   &FshomeRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}

func NewTestRepos() Repos {
	return Repos{
		Fs:       &FsRepository{}, // todo
		Fshome:   &FshomeMockRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}
