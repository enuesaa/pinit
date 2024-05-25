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
		Short: "Initialize pinit app. Run database migration and setup application.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := usecase.EnvCheck(repos); err != nil {
				return err
			}
			if err := usecase.DBCreate(repos); err != nil {
				return err
			}
			if err := usecase.DBOpen(repos); err != nil {
				return err
			}
			if err := usecase.DBMigrate(repos); err != nil {
				return err
			}
			fmt.Printf("succeed\n")

			return nil
		},
	}

	return cmd
}
