package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logstay/project-chassi/internal/domain"
	"github.com/sirupsen/logrus"
)

func (s *server) obterExemplo(c *gin.Context) {

	resp, err := s.endpoint.ObterExemploEndpoint(c, nil)
	if err != nil {
		logrus.Error(err)
		s.ResponseErrorMessage(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(resp.(domain.CustomerResponse).Code, resp.(domain.CustomerResponse).Response)
}

func (s *server) inserirExemplo(c *gin.Context) {
	var Exemplo domain.Exemplo

	if err := c.ShouldBind(&Exemplo); err != nil {
		logrus.Error(err)
		s.ResponseErrorMessage(c, http.StatusBadRequest, err.Error())
	}

	resp, err := s.endpoint.InserirExemploEndpoint(c, Exemplo)
	if err != nil {
		logrus.Error(err)
		s.ResponseErrorMessage(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(resp.(domain.CustomerResponse).Code, resp.(domain.CustomerResponse).Response)
}
