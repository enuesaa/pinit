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
			return usecase.OpenDb(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CloseDb(repos)
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
