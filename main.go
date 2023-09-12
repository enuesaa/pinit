package main

import (
	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "pinit",
		Short: "A CLI client for a personal note-taking application.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// setup repository
	repos := repository.NewRepos()
	configSrv := service.NewConfigSevice(repos)
	databaseDsn, err := configSrv.ReadDatabaseDsn()
	if err == nil {
		repos.Database.WithDsn(databaseDsn)
	}

	// sub commands
	cmd.AddCommand(cli.CreateConfigureCmd(repos))
	cmd.AddCommand(cli.CreateNewCmd(repos))
	cmd.AddCommand(cli.CreateLsCmd(repos))
	cmd.AddCommand(cli.CreateDescribeCmd(repos))
	cmd.AddCommand(cli.CreateRmCmd(repos))

	// disable default
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.Execute()
}
