package repository

type FshomeMockRepository struct {
	Configfile string
}

func (repo *FshomeMockRepository) isFileOrDirExist(path string) bool {
	return true
}

func (repo *FshomeMockRepository) homedir() (string, error) {
	return "/", nil
}

func (repo *FshomeMockRepository) IsRegistryExist(registryName string) bool {
	return true
}

func (repo *FshomeMockRepository) CreateRegistry(registryName string) error {
	return nil
}

func (repo *FshomeMockRepository) IsFileExist(registryName string, path string) bool {
	return true
}

func (repo *FshomeMockRepository) WriteFile(registryName string, path string, content string) error {
	return nil
}

func (repo *FshomeMockRepository) ReadFile(registryName string, path string) (string, error) {
	return repo.Configfile, nil
}
