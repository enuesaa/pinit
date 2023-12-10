package repository

type ConfigMockRepository struct {
	items map[string]string
}

func (repo *ConfigMockRepository) Init() {}

func (repo *ConfigMockRepository) Load() error {
	return nil
}

func (repo *ConfigMockRepository) DbHost() string {
	return repo.items["dbhost"]
}

func (repo *ConfigMockRepository) SetDbHost(dbhost string) {
	repo.items["dbhost"] = dbhost
}

func (repo *ConfigMockRepository) DbUsername() string {
	return repo.items["dbusername"]
}

func (repo *ConfigMockRepository) SetDbUsername(dbusername string) {
	repo.items["dbusername"] = dbusername
}

func (repo *ConfigMockRepository) DbPassword() string {
	return repo.items["dbpassword"]
}

func (repo *ConfigMockRepository) SetDbPassword(dbpassword string) {
	repo.items["dbpassword"] = dbpassword
}

func (repo *ConfigMockRepository) DbName() string {
	return repo.items["dbname"]
}

func (repo *ConfigMockRepository) SetDbName(dbname string) {
	repo.items["dbname"] = dbname
}

func (repo *ConfigMockRepository) ChatgptToken() string {
	return repo.items["chatgpttoken"]
}

func (repo *ConfigMockRepository) SetChatgptToken(chatgpttoken string) {
	repo.items["chatgpttoken"] = chatgpttoken
}

func (repo *ConfigMockRepository) Write() error {
	return nil
}
