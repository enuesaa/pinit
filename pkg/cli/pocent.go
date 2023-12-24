package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/enuesaa/pinit/ent"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePocentsCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pocent",
		Short: "pocent",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("poc ent command\n")

			dsn := repos.Database.Dsn()
			client, err := ent.Open("mysql", dsn)
			if err != nil {
				log.Fatalf("Error: failed to open\n%s\n", err.Error())
			}
			defer client.Close()

			binder, err := client.Binder.Create().SetName("hello").SetCategory("test").Save(context.Background())
			if err != nil {
				log.Fatalf("Error: failed to create binder\n%s\n", err.Error())
			}
			fmt.Printf("%+v\n", binder)
		},
	}

	return cmd
}