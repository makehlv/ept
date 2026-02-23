package swagger

import (
	"log/slog"

	"github.com/makehlv/ept/clients"
	"github.com/makehlv/ept/config"
	"github.com/makehlv/ept/repositories"
)

type SwaggerService struct {
	logger  *slog.Logger
	clients *clients.Clients
	config  *config.Config

	repositories *repositories.Repositories
}

func NewSwaggerService(clients *clients.Clients, logger *slog.Logger, config *config.Config, repositories *repositories.Repositories) *SwaggerService {
	return &SwaggerService{clients: clients, logger: logger, config: config, repositories: repositories}
}
