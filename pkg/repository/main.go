package repository

type Repos struct {
	Fshome   FshomeRepositoryInterface
	Database DatabaseRepositoryInterface
	Prompt   PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Fshome:   &FshomeRepository{},
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}

func NewTestRepos() Repos {
	fshome := &FshomeMockRepository{
		Configfile: "",
	}

	return Repos{
		Fshome:   fshome,
		Database: &DatabaseRepository{},
		Prompt:   &Prompt{},
	}
}
