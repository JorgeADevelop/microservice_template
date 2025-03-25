package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"www.marawa.com/microservice_service/internal/infra/config"
	fx_module "www.marawa.com/microservice_service/internal/infra/fx"
	"www.marawa.com/microservice_service/pkg/translater"
	"www.marawa.com/microservice_service/pkg/validation"
)

func main() {
	app := fx.New(
		fx.Invoke(translater.NewTranslater),
		fx.Invoke(validation.NewValidator),
		fx.Provide(
			config.NewConfig,
		),
		fx_module.DBModule,
		fx_module.HTTPModule,
		fx_module.RouterModule,
		fx.Invoke(func(*gin.Engine) {}),
	)
	app.Run()
}
