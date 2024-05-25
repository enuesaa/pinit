package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateInitCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize pinit. Run database migration and setup pinit",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := usecase.EnvCheck(repos); err != nil {
				return err
			}

			if !usecase.DBIsExist(repos) {
				fmt.Printf("database migration start\n")
				if err := usecase.DBSetup(repos); err != nil {
					return err
				}
				fmt.Printf("database migration succeeded\n")
			}

			if err := usecase.DBOpen(repos); err != nil {
				return err
			}
			defer usecase.DBClose(repos)

			fmt.Printf("app setup..\n")
			if err := usecase.ConfigAsk(repos); err != nil {
				return err
			}
			fmt.Printf("succeeded\n")

			return nil
		},
	}

	return cmd
}
