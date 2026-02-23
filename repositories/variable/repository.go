package variable

import (
	"log/slog"

	"github.com/makehlv/ept/config"
)

type VariableRepository struct {
	logger  *slog.Logger
	config  *config.Config
}

func NewVariableRepository(logger *slog.Logger, config *config.Config) *VariableRepository {
	return &VariableRepository{logger: logger, config: config}
}
