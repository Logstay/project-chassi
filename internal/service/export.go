package service

import (
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
	Exemplo "github.com/logstay/project-church-service/internal/service/exemplo"
)

type ServiceFactory interface {
	Exemplo() Exemplo.Service
}

type serviceFactory struct {
	Exemploervice Exemplo.Service
}

func NewServiceFactory(db *sqlx.DB, logger log.Logger) ServiceFactory {
	return &serviceFactory{
		Exemploervice: Exemplo.NewService(db, logger),
	}
}

func (sf *serviceFactory) Exemplo() Exemplo.Service {
	return sf.Exemploervice
}
