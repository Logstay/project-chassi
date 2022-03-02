package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log/level"
	"github.com/logstay/project-church-service/internal/domain"
	"github.com/sirupsen/logrus"
)

func (s *server) obterExemplo(c *gin.Context) {

	resp, err := s.endpoint.ObterExemploEndpoint(c, nil)
	if err != nil {
		logrus.Error(err)
	}

	c.JSON(resp.(domain.CustomerResponse).Code, resp.(domain.CustomerResponse).Response)

}

func (s *server) inserirExemplo(c *gin.Context) {
	var Exemplo domain.Exemplo

	if err := c.ShouldBind(&Exemplo); err != nil {
		_ = level.Error(*s.logger).Log("message", "invalid request")
	}

	resp, err := s.endpoint.InserirExemploEndpoint(c, Exemplo)
	if err != nil {
		logrus.Error(err)
	}

	c.JSON(resp.(domain.CustomerResponse).Code, resp.(domain.CustomerResponse).Response)

}
