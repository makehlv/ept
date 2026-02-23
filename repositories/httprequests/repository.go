package httprequests

import (
	"log/slog"

	"github.com/makehlv/ept/config"
)

type HttpRequestRepository struct {
	logger  *slog.Logger
	config  *config.Config
}

func NewHttpRequestRepository(logger *slog.Logger, config *config.Config) *HttpRequestRepository {
	return &HttpRequestRepository{logger: logger, config: config}
}
