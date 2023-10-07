package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateSetupMigrationCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "migration",
		Short: "migration tables.",
		Run: func(cmd *cobra.Command, args []string) {
			databaseDsn, err := cmd.Flags().GetString("database-dsn")
			if databaseDsn == "" || err != nil {
				fmt.Printf("error: Missing required flag `--database-dsn`.\n")
				return
			}
			repos.Database.WithDsn(databaseDsn)
			fmt.Println("this is migration task.")
			isTableExist, err := repos.Database.IsTableExist(&service.Note{})

			if err != nil {
				fmt.Printf("error: %s\n", err)
				return
			}
			fmt.Printf("isTableExist: %t\n", isTableExist)
			if !isTableExist {
				repos.Database.CreateTable(&service.Note{})
				fmt.Printf("note table created.\n")
				repos.Database.CreateTable(&service.Tag{})
				fmt.Printf("tag table created.\n")
			}
		},
	}
	cmd.Flags().String("database-dsn", "", "[Required] Database Dsn")

	return cmd
}
