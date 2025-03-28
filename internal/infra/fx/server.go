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
			go func() {
				if serveErr := srv.Serve(ln); serveErr != http.ErrServerClosed {
					fmt.Printf("Error while starting server: %v\n", serveErr)
				}
			}()
			return nil

		},
		OnStop: func(ctx context.Context) error {
			if err := srv.Shutdown(ctx); err != nil {
				fmt.Printf("Error while shutting down server: %v\n", err)
				return err
			}
			return nil
		},
	})

	return router
}

var HTTPModule = fx.Options(
	fx.Provide(StartServer),
)
