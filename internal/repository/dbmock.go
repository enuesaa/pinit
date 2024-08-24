package repository

import (
	"github.com/enuesaa/pinit/internal/ent"
)

type DatabaseMockRepository struct{}

func (repo *DatabaseMockRepository) CheckEnv() error {
	return nil
}

func (repo *DatabaseMockRepository) IsDBExist() bool {
	return true
}

func (repo *DatabaseMockRepository) CreateDB() error {
	return nil
}

func (repo *DatabaseMockRepository) Open() error {
	return nil
}

func (repo *DatabaseMockRepository) Close() error {
	return nil
}

func (repo *DatabaseMockRepository) Db() (*ent.Client, error) {
	return nil, nil
}

func (repo *DatabaseMockRepository) Migrate() error {
	return nil
}

func (repo *DatabaseMockRepository) Appconf() *ent.AppconfClient {
	return nil
}

func (repo *DatabaseMockRepository) Binder() *ent.BinderClient {
	return nil
}

func (repo *DatabaseMockRepository) Note() *ent.NoteClient {
	return nil
}

func (repo *DatabaseMockRepository) Action() *ent.ActionClient {
	return nil
}

// db.Use(hook.On(func(next ent.Mutator) ent.Mutator {
// 	return hook.BinderFunc(func(ctx context.Context, m *ent.BinderMutation) (ent.Value, error) {
// 		fmt.Printf("a\n")
// 		return next.Mutate(ctx, m)
// 	})
// }, ent.OpCreate))

// db.Intercept(
// 	ent.InterceptFunc(func(next ent.Querier) ent.Querier {
// 		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
// 			value, err := next.Query(ctx, query)
// 			fmt.Printf("selected: %+v\n", value)
// 			return value, err
// 		})
// 	}),
// )
