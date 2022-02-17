package cmd

import (
	"github.com/brianbi/gin-starter/bootstrap"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Start Http REST API",
		Run:   runServer,
	}
)

func runServer(cmd *cobra.Command, args []string) {
	fx.New(
		fx.Provide(
			bootstrap.NewLogger,
		),
		fx.Invoke(

		),
	).Run()
}


