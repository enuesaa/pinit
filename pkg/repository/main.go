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

func NewTestRepos(dbHost string) Repos {
	database := &DatabaseRepository{
		DbHost: dbHost,
		DbUsername: "test",
		DbPassword: "test",
		DbName: "test_pinit",
	}

	return Repos{
		Fshome:   &FshomeMockRepository{},
		Database: database,
		Prompt:   &Prompt{},
	}
}
