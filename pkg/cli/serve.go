package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/serve"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch the web console",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := usecase.ConfigInit(repos); err != nil {
				return err
			}
			return usecase.OpenDbConnection(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CloseDbConnection(repos)
		},
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")

			serveCtl := serve.NewServeCtl(repos)
			serveCtl.Serve(port)
		},
	}
	cmd.Flags().Int("port", 3000, "port. Default: 3000")

	return cmd
}
