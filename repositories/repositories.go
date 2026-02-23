package repositories

import (
	"log/slog"

	"github.com/makehlv/ept/config"
	"github.com/makehlv/ept/repositories/httprequests"
	"github.com/makehlv/ept/repositories/swagger"
	"github.com/makehlv/ept/repositories/variable"
)

type Repositories struct {
	Swagger     *swagger.SwaggerRepository
	Variable    *variable.VariableRepository
	HttpRequest *httprequests.HttpRequestRepository
}

func NewRepositories(logger *slog.Logger, config *config.Config) *Repositories {
	return &Repositories{
		Swagger:     swagger.NewSwaggerRepository(logger, config),
		Variable:    variable.NewVariableRepository(logger, config),
		HttpRequest: httprequests.NewHttpRequestRepository(logger, config),
	}
}
