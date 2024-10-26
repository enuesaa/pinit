package repository

type DatabaseMockRepository struct {
	DbRepository
}

func (repo *DatabaseMockRepository) IsDBExist() bool {
	return true
}
