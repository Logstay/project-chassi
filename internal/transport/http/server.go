package transport

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"

	"github.com/logstay/project-chassi/internal/endpoint"
)

type server struct {
	endpoint *endpoint.Endpoints
	logger   *log.Logger
}

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(context context.Context, endpoint *endpoint.Endpoints, logger *log.Logger) http.Handler {
	rest := &server{
		endpoint: endpoint,
		logger:   logger,
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", rest.HealthCheckHandler)

	Exemplo := r.Group("/v1/exemplo")
	{
		Exemplo.GET("", rest.obterExemplo)
		Exemplo.POST("", rest.inserirExemplo)
	}

	err := r.Run(":8080")
	logrus.Error(err)

	return r
}

func (s *server) ResponseErrorMessage(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, struct {
		Message string `json:"message"`
	}{message})
}
