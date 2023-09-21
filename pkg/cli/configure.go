package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
)

func CreateConfigureCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "configure",
		Run: func(cmd *cobra.Command, args []string) {
			isReadTask, _ := cmd.Flags().GetBool("read")
			isMigrationTask, _ := cmd.Flags().GetBool("migration")
			passedDatabaseDsn, _ := cmd.Flags().GetString("database-dsn")
			configSrv := service.NewConfigSevice(repos)

			if isReadTask {
				config, err := configSrv.Read()
				if err != nil {
					return
				}
				fmt.Printf("Database DSN: %s \n", config.DatabaseDsn)
				return
			}

			// TODO: refactor
			if isMigrationTask {
				if passedDatabaseDsn == "" {
					fmt.Println("please provide database-dsn flag.")
					return
				}
				repos.Database.WithDsn(passedDatabaseDsn)
				fmt.Println("this is migration task.")
				isTableExist, err := repos.Database.IsTableExist(&service.Note{})

				//root@tcp(localhost)/testpinit?interpolateParams=true&parseTime=true

				if err != nil {
					fmt.Printf("error: %s\n", err)
					return
				}
				fmt.Printf("isTableExist: %t\n", isTableExist)
				if !isTableExist {
					repos.Database.CreateTable(&service.Note{})
					fmt.Printf("table created.\n")
				}
				return
			}

			databaseDsn, err := textinput.New("Database DSN").RunPrompt()
			if err != nil {
				return
			}
			if err := configSrv.WriteDatabaseDsn(databaseDsn); err != nil {
				fmt.Println(err)
				return
			}			
		},
	}
	cmd.Flags().Bool("read", false, "Read Config")
	cmd.Flags().Bool("migration", false, "Exuecute database migration if notes table does not exist.")
	cmd.Flags().String("database-dsn", "", "Using when migration flag provided.")

	return cmd
}
