package repository

type ConfigMockRepository struct {
	ConfigRepository
}

func (repo *ConfigMockRepository) Init() {}

func (repo *ConfigMockRepository) Load() error {
	return nil
}

func (repo *ConfigMockRepository) Write() error {
	return nil
}
