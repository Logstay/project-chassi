package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/logstay/project-church-service/internal/domain"
	"github.com/logstay/project-church-service/internal/service"
)

func makeObterExemploEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		resp, err := s.Exemplo().ObterExemplo(ctx)
		if err != nil {
			_ = level.Error(logger).Log("message", "invalid request")
			return domain.CustomerResponse{
				Code:     http.StatusBadRequest,
				Response: err.Error(),
			}, nil
		}

		_ = level.Error(logger).Log("message", "ok")

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
			_ = level.Error(logger).Log("message", "invalid request")
			return domain.CustomerResponse{
				Code:     http.StatusBadRequest,
				Response: err.Error(),
			}, nil
		}

		_ = level.Error(logger).Log("message", "ok")

		return domain.CustomerResponse{
			Code: http.StatusOK,
		}, nil
	}
}
