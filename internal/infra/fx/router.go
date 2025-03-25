package fx

import (
	"go.uber.org/fx"
	"www.marawa.com/microservice_service/internal/infra/routers"
)

var RouterModule = fx.Options(
	fx.Invoke(routers.NewHealthRouter),
)
