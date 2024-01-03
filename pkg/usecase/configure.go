package usecase

import (
	"github.com/enuesaa/pinit/pkg/repository"
)

func Configure(repos repository.Repos) error {
	dbHost, err := repos.Prompt.Ask("DB Host", repos.Config.DbHost())
	if err != nil {
		return err
	}
	repos.Config.SetDbHost(dbHost)
	dbUsername, err := repos.Prompt.Ask("DB Username", repos.Config.DbUsername())
	if err != nil {
		return err
	}
	repos.Config.SetDbUsername(dbUsername)
	dbPassword, err := repos.Prompt.Ask("DB Password", repos.Config.DbPassword())
	if err != nil {
		return err
	}
	repos.Config.SetDbPassword(dbPassword)
	dbName, err := repos.Prompt.Ask("DB Name", repos.Config.DbName())
	if err != nil {
		return err
	}
	repos.Config.SetDbName(dbName)

	chatgptToken, err := repos.Prompt.Ask("OpenAI API Token", repos.Config.ChatgptToken())
	if err != nil {
		return err
	}
	repos.Config.SetChatgptToken(chatgptToken)

	return repos.Config.Write()
}
