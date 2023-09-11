package repository

type Repos struct {
	Fshome FshomeRepositoryInterface
	Database DatabaseRepositoryInterface
}

func NewRepos() Repos {
	return Repos {
		Fshome: &FshomeRepository{},
		Database: &DatabaseRepository{},
	}
}
