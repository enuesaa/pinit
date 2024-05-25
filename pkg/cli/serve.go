package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve pinit",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.DBOpen(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.DBClose(repos)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			ctl := usecase.NewServeCtl(repos, port)
			if err := ctl.Open(); err != nil {
				// ignore this err because this is not critical for app.
				repos.Log.Info("failed to open url because `%s`", err.Error())
			}

			return ctl.Serve()
		},
	}
	cmd.Flags().Int("port", 3000, "port. Default: 3000")

	return cmd
}
