package server

import (
	"fmt"
	"github.com/brianbi/gin-starter/config"
	"go.uber.org/fx"
)

var Provides = fx.Options(
	fx.Provide(config.NewConfiguration),
	fx.Invoke(initConfig),
)

func initConfig(config *config.Configuration)  {
	fmt.Println(config.GetString("app.name"))
}