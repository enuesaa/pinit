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
