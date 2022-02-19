package server

import (
	"github.com/brianbi/gin-starter/bootstrap/server"
	"github.com/brianbi/gin-starter/config"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	configFile string
	HttpServerCmd = &cobra.Command{
		Use:          "http-server",
		Short:        "Start Http server",
		Example:      "{execfile} server -c config/config.yaml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(configFile) > 0 {
				config.SetConfigPath(configFile)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}
)

func init() {
	pf := HttpServerCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c", "./config/config.yaml", "this parameter is used to start the service application")
	//if err := cobra.MarkFlagRequired(pf, "config"); err != nil {
	//	panic(err.Error())
	//}
}

func runServer() {
	fx.New(server.Provides, fx.NopLogger).Run()
}
