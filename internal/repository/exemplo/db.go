package exemplo

import (
	"github.com/jmoiron/sqlx"
	"github.com/logstay/project-church-service/internal/domain"
	rputils "github.com/logstay/project-church-service/internal/repository/rp_utils"
)

type ExemploRepository interface {
	ObterExemplo() ([]domain.Exemplo, error)

	InserirExemplo(exemplo domain.Exemplo) (int64, error)
}

type exemploRepository struct {
	db *sqlx.DB
}

func NewExemploRepository(db *sqlx.DB) ExemploRepository {
	return &exemploRepository{
		db: db,
	}
}

func (ep *exemploRepository) ObterExemplo() ([]domain.Exemplo, error) {

	Exemplo := []domain.Exemplo{}

	err := ep.db.Select(&Exemplo, "SELECT * FROM exemplo")

	return Exemplo, err
}

func (ep *exemploRepository) InserirExemplo(exemplo domain.Exemplo) (int64, error) {
	var resp *sqlx.Rows
	var err error
	var id int64

	err = rputils.Transaction(ep.db, func(t *sqlx.Tx) error {
		resp, err = ep.db.NamedQuery(qryInserirExemplo, exemplo)

		id, err = rputils.CloseDBTransactionReturningID(resp, err)

		return err
	})

	return id, err
}
