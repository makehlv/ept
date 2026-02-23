package variable

import (
	"log/slog"

	"github.com/makehlv/ept/clients"
	"github.com/makehlv/ept/config"
	"github.com/makehlv/ept/repositories"
)

type VariableService struct {
	logger       *slog.Logger
	clients      *clients.Clients
	config       *config.Config
	repositories *repositories.Repositories
}

func NewVariableService(
	clients *clients.Clients, logger *slog.Logger, config *config.Config, repositories *repositories.Repositories) *VariableService {
	return &VariableService{clients: clients, logger: logger, config: config, repositories: repositories}
}
