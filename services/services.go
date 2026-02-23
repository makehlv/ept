package services

import (
	"log/slog"

	"github.com/makehlv/ept/clients"
	"github.com/makehlv/ept/config"
	"github.com/makehlv/ept/repositories"
	"github.com/makehlv/ept/services/swagger"
	"github.com/makehlv/ept/services/variable"
)

type Services struct {
	Swagger  *swagger.SwaggerService
	Variable *variable.VariableService
}

func NewServices(clients *clients.Clients, logger *slog.Logger, config *config.Config, repositories *repositories.Repositories) *Services {
	return &Services{
		Swagger:  swagger.NewSwaggerService(clients, logger, config, repositories),
		Variable: variable.NewVariableService(clients, logger, config, repositories),
	}
}
