package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/logstay/project-church-service/internal/service"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	Health                 endpoint.Endpoint
	ObterExemploEndpoint   endpoint.Endpoint
	InserirExemploEndpoint endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s service.ServiceFactory, logger log.Logger) Endpoints {
	return Endpoints{
		Health:                 makeHealthEndpoint(s, logger),
		ObterExemploEndpoint:   makeObterExemploEndpoint(s, logger),
		InserirExemploEndpoint: makeInserirExemploEndpoint(s, logger),
	}
}
