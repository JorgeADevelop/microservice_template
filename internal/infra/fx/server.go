package fx

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"www.marawa.com/microservice_service/internal/infra/config"
)

func StartServer(
	lc fx.Lifecycle,
	cfg *config.Config,
) *gin.Engine {
	router := gin.Default()

	srv := &http.Server{Addr: fmt.Sprintf("%s:%s", cfg.AppConfig.Host, cfg.AppConfig.Port), Handler: router}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			go srv.Serve(ln)
			return nil

		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)
			return nil
		},
	})

	return router
}

// HTTPModule es el módulo FX para configurar el servidor HTTP
var HTTPModule = fx.Options(
	fx.Provide(StartServer),
)
