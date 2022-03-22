package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/logstay/project-chassi/internal/domain"
	"github.com/logstay/project-chassi/internal/service"
	"github.com/sirupsen/logrus"
)

// makeHealthEndpoint return if service up
func makeHealthEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		logrus.WithFields(logrus.Fields{
			"fluxo":    "ok",
			"endpoint": "makeHealthEndpoint",
		})

		return domain.CustomerResponse{
			Code:     http.StatusOK,
			Response: "Service project UP",
		}, nil
	}
}
