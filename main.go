package main

import (
	"github.com/spf13/cobra"
)

func main() {
	cli := &cobra.Command{
		Use:  "pinit",
		Args: cobra.MinimumNArgs(0),
		Run:  func(cmd *cobra.Command, args []string) {},
	}

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true

	cli.Execute()
}
