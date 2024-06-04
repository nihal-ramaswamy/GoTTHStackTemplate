package main

import (
	"github.com/gin-gonic/gin"
	serverconfig "github.com/nihal-ramaswamy/gotta_stack/internal/config/server"
	"github.com/nihal-ramaswamy/gotta_stack/internal/utils"
	fx_utils "github.com/nihal-ramaswamy/gotta_stack/internal/utils/fx"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(utils.NewProduction),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		fx_utils.ConfigModule,
		fx_utils.MicroServicesModule,

		fx.Invoke(Invoke),
	).Run()
}

func Invoke(server *gin.Engine, config *serverconfig.Config, log *zap.Logger) {
	err := server.Run(config.Port)
	if nil != err {
		log.Error(err.Error())
	}
}
