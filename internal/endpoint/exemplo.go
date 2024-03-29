package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/logstay/project-chassi/internal/domain"
	"github.com/logstay/project-chassi/internal/service"
	"github.com/sirupsen/logrus"
)

func makeObterExemploEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		resp, err := s.Exemplo().ObterExemplo(ctx)
		if err != nil {
			logrus.Debug("Error flux -> ", err)
			return domain.CustomerResponse{
				Code:     http.StatusBadRequest,
				Response: err.Error(),
			}, nil
		}

		logrus.WithFields(logrus.Fields{
			"fluxo":    "ok",
			"endpoint": "makeObterExemploEndpoint",
		})

		return domain.CustomerResponse{
			Code:     http.StatusOK,
			Response: resp,
		}, nil
	}
}

func makeInserirExemploEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if request == nil {
			return domain.CustomerResponse{
				Code:     http.StatusBadRequest,
				Response: errors.New("requisição sem parâmetro"),
			}, nil
		}

		Exemplo := request.(domain.Exemplo)
		err := s.Exemplo().AdicionarExemplo(ctx, Exemplo)
		if err != nil {
			logrus.Debug("Error flux -> ", err)
			return domain.CustomerResponse{
				Code:     http.StatusBadRequest,
				Response: err.Error(),
			}, nil
		}

		logrus.WithFields(logrus.Fields{
			"fluxo":    "ok",
			"endpoint": "makeInserirExemploEndpoint",
		})

		return domain.CustomerResponse{
			Code: http.StatusOK,
		}, nil
	}
}
