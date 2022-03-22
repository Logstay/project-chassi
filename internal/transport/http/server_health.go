package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logstay/project-chassi/internal/domain"
	"github.com/sirupsen/logrus"
)

func (s *server) HealthCheckHandler(c *gin.Context) {

	resp, err := s.endpoint.Health(c, nil)
	if err != nil {
		logrus.Error(err)
		s.ResponseErrorMessage(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(resp.(domain.CustomerResponse).Code, resp.(domain.CustomerResponse).Response)
}
