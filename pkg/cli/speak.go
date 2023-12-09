package cli

import (
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateSpeakCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "speak",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repository.NewRepos())
			config, err := configSrv.Read()
			if err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}

			aiSrv := service.NewAiService(repos)
			if err := aiSrv.Speak(config.ChatgptToken); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
		},
	}

	return cmd
}
