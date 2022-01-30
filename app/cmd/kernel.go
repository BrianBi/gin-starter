package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	Version = "1.0.0"

	rootCmd = &cobra.Command{
		Use:     "gin-starter",
		Version: Version,
		Short:   "gin-starter Management CLI",
		Run: func(cmd *cobra.Command, args []string) {
			httpCmd.Run(cmd, args)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(httpCmd)
}