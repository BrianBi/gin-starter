package cmd

import (
	"errors"
	"github.com/brianbi/gin-starter/cmd/server"
	"github.com/spf13/cobra"
	"os"
)

var (
	Version = "1.0.0"

	rootCmd = &cobra.Command{
		Use:    "gin-starter",
		Short:  "gin-starter Management CLI",
		Version:	Version,
		SilenceUsage: true,
		Long:	`gin-admin`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New(
					"requires at least one arg, you can view the available parameters through `--help`",
				)
			}
			return nil
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Run: func(cmd *cobra.Command, args []string) {
			// pass
		},
	}

)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(server.HttpServerCmd)
}

