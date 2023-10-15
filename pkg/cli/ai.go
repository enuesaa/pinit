package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateAiCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "ai",
		Short: "call openai api",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.Init(); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			config, err := configSrv.Read()
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			chatgptSrv := service.NewAiService(repos)
			res, err := chatgptSrv.Call(config.ChatgptToken)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			fmt.Printf("res: %s\n", res)
		},
	}

	return cmd
}