package mocks

import (
	"github.com/logstay/project-church-service/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ExemploMock struct {
	mock.Mock
}

func (m *ExemploMock) ObterExemplo() ([]domain.Exemplo, error) {
	args := m.Called()

	return args.Get(0).([]domain.Exemplo), args.Error(1)
}

func (m *ExemploMock) InserirExemplo(Exemplo domain.Exemplo) (int64, error) {
	args := m.Called()

	return args.Get(0).(int64), args.Error(1)
}
