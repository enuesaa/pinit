package repository

type Repos struct {
	Config ConfigRepositoryInterface
	Database DatabaseRepositoryInterface
}

func NewRepos() Repos {
	return Repos {
		Config: &ConfigRepository{},
		Database: &DatabaseRepository{},
	}
}
