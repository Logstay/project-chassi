package Exemplo

import (
	"context"
	"errors"
	"testing"

	"github.com/logstay/project-church-service/internal/domain"
	"github.com/logstay/project-church-service/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestObterExemplo(t *testing.T) {
	var ctx context.Context
	ExemploMock := &mocks.ExemploMock{}
	service := &service{
		ExemploRepository: ExemploMock,
	}

	t.Run("Não deveria obter pessosa caso vazio", func(t *testing.T) {

		ExemploMock.On("ObterExemplo").Times(1).Return([]domain.Exemplo{{}}, nil)

		resp, err := service.ObterExemplo(ctx)
		assert.Nil(t, err)
		assert.Equal(t, 0, int(resp[0].ID.Int64))
		ExemploMock.AssertExpectations(t)
	})

	t.Run("Deveria obter Exemplo caso sucesso", func(t *testing.T) {

		ExemploMock.On("ObterExemplo").Times(1).Return([]domain.Exemplo{
			{
				ID: null.IntFrom(1),
			},
		}, nil)

		resp, err := service.ObterExemplo(ctx)
		assert.Nil(t, err)
		assert.Equal(t, null.IntFrom(1), resp[0].ID)
		ExemploMock.AssertExpectations(t)
	})

	t.Run("Não deveria obter Exemplo caso falha", func(t *testing.T) {

		ExemploMock.On("ObterExemplo").Times(1).Return([]domain.Exemplo{}, errors.New("erro"))

		resp, err := service.ObterExemplo(ctx)
		assert.EqualError(t, err, "Não foi possível obter Exemplo")
		assert.Len(t, resp, 0)
		ExemploMock.AssertExpectations(t)
	})

}

func TestInserirExemplo(t *testing.T) {
	var ctx context.Context
	ExemploMock := &mocks.ExemploMock{}
	service := &service{
		ExemploRepository: ExemploMock,
	}

	t.Run("Deveria inserir Exemplo caso sucesso", func(t *testing.T) {

		ExemploMock.On("InserirExemplo").Times(1).Return(int64(1), nil)

		err := service.AdicionarExemplo(ctx, domain.Exemplo{
			ID: null.IntFrom(1),
		})

		assert.Nil(t, err)
		ExemploMock.AssertExpectations(t)
	})

	t.Run("Não deveria inserir Exemplo caso sucesso", func(t *testing.T) {

		ExemploMock.On("InserirExemplo").Times(1).Return(int64(0), errors.New("erro"))

		err := service.AdicionarExemplo(ctx, domain.Exemplo{
			ID: null.IntFrom(1),
		})

		assert.EqualError(t, err, "Não foi possível inserir Exemplo")
		ExemploMock.AssertExpectations(t)
	})

}
