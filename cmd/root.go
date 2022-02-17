package cmd

import (
	"github.com/brianbi/gin-starter/config"
	"github.com/spf13/cobra"
	"os"
)

var (
	Version = "1.0.0"

	rootCmd = &cobra.Command{
		Use:     "go-starter",
		Version: Version,
		Short:   "go-starter Management CLI",
		Run: func(cmd *cobra.Command, args []string) {
			httpCmd.Run(cmd, args)
		},
	}

	configFile string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(httpCmd)
}

func initConfig() {
	config.Initialize(configFile)
}
