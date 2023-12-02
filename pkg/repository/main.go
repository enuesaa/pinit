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

func NewTestRepos(databseDsn string) Repos {
	database := &DatabaseRepository{}
	database.WithDsn(databseDsn)

	return Repos{
		Fshome:   &FshomeMockRepository{},
		Database: database,
		Prompt:   &Prompt{},
	}
}
