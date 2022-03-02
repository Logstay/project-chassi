package Exemplo

import (
	"context"
	"errors"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jmoiron/sqlx"
	"github.com/logstay/project-church-service/internal/domain"
	ExemploRepository "github.com/logstay/project-church-service/internal/repository/exemplo"
)

type Service interface {
	ObterExemplo(ctx context.Context) ([]domain.Exemplo, error)

	AdicionarExemplo(ctx context.Context, Exemplo domain.Exemplo) error
}

type service struct {
	ExemploRepository ExemploRepository.ExemploRepository
	logger            log.Logger
}

func NewService(db *sqlx.DB, logger log.Logger) *service {
	return &service{
		ExemploRepository: ExemploRepository.NewExemploRepository(db),
		logger:            logger,
	}
}

func (s *service) ObterExemplo(ctx context.Context) ([]domain.Exemplo, error) {
	logger := log.With(s.logger, "metodo", "ObterExemplo")

	Exemplo, err := s.ExemploRepository.ObterExemplo()
	if err != nil {
		level.Error(logger)

		return []domain.Exemplo{}, errors.New("Não foi possível obter Exemplo")
	}

	return Exemplo, nil
}

func (s *service) AdicionarExemplo(ctx context.Context, Exemplo domain.Exemplo) error {

	_, err := s.ExemploRepository.InserirExemplo(Exemplo)
	if err != nil {
		return errors.New("Não foi possível inserir Exemplo")
	}

	return nil
}
