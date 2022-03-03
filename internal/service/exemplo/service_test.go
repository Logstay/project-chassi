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

		resp, err := service.ObterExemplo(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		ExemploMock.AssertExpectations(t)
	})

	t.Run("Deveria obter Exemplo caso sucesso", func(t *testing.T) {

		resp, err := service.ObterExemplo(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
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
