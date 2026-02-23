package swagger

import (
	"log/slog"

	"github.com/makehlv/ept/config"
)

type SwaggerRepository struct {
	logger  *slog.Logger
	config  *config.Config
}

func NewSwaggerRepository(logger *slog.Logger, config *config.Config) *SwaggerRepository {
	return &SwaggerRepository{logger: logger, config: config}
}
