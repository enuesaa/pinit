package repository

type Env struct {
	DATABASE_DSN string
}

type Repos struct {
	Database DatabaseRepositoryInterface
}

func NewRepos(env Env) Repos {
	return Repos{
		Database: &DatabaseRepository{
			Dsn: env.DATABASE_DSN,
		},
	}
}
