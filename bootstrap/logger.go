package bootstrap

//import (
//	"context"
//	"fmt"
//	"github.com/brianbi/gin-starter/config"
//	"github.com/brianbi/gin-starter/pkg/env"
//	"github.com/brianbi/gin-starter/pkg/logger"
//	"go.uber.org/fx"
//	"go.uber.org/zap"
//)
//
//// NewLogger logger 构造函数
//func NewLogger(lc fx.Lifecycle) (*zap.Logger, error) {
//	loggers, err := logger.NewJSONLogger(
//		logger.WithField("domain", fmt.Sprintf("%s[%s]", config.Get("app.name"), env.Active().Value())),
//		logger.WithTimeLayout(config.GetString("log.time_layout")),
//		logger.WithFileRotationP(
//			config.GetString("log.filename"),
//			config.GetInt("log.max_size"),
//			config.GetInt("log.max_backup"),
//			config.GetInt("log.max_age"),
//			config.GetBool("log.compress"),
//			config.GetString("log.type"),
//		),
//	)
//
//	if err != nil {
//		panic(err)
//	}
//
//	lc.Append(fx.Hook{
//		OnStart: func(i context.Context) error {
//			loggers.Info("logger on start...")
//			return nil
//		},
//		OnStop: func(i context.Context) error {
//			loggers.Info("logger on stop...")
//			return nil
//		},
//	})
//
//	defer func() {
//		_ = loggers.Sync()
//	}()
//
//	return loggers, nil
//}
