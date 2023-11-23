package main

import (
	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:     "pinit",
		Short:   "A personal note-taking app",
		Version: "0.0.3",
	}

	repos := repository.NewRepos()
	cmd.AddCommand(cli.CreateConfigureCmd(repos))
	cmd.AddCommand(cli.CreateLsCmd(repos))
	cmd.AddCommand(cli.CreateNewCmd(repos))
	cmd.AddCommand(cli.CreateRmCmd(repos))
	cmd.AddCommand(cli.CreateUpCmd(repos))
	usecase.Init(repos)

	// disable default
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.Execute()
}
