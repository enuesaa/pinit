package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateSetupChatgptCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "chatgpt",
		Short: "call chatgpt api.",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")
			res, err := service.CallChatgptApi(token)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			fmt.Printf("res: %s\n", res)
		},
	}
	cmd.Flags().String("token", "", "token")
	cmd.MarkFlagRequired("token")

	return cmd
}
