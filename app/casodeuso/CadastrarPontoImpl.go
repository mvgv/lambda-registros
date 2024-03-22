package casodeuso

import (
	"github.com/mvgv/lambda-registros/app/dominio"
	"github.com/mvgv/lambda-registros/app/infraestrutura/repositorio"
)

type CadastrarPontoImpl struct {
	pontoRepo repositorio.PontoRepositorio
}

func NewCadastrarPontoImpl(pontoRepo repositorio.PontoRepositorio) *CadastrarPontoImpl {
	return &CadastrarPontoImpl{
		pontoRepo: pontoRepo,
	}
}

func (c *CadastrarPontoImpl) CadastrarPontoDoDia(ponto dominio.Ponto, email string) (string, error) {

	/*registrosDoDia, err := c.pontoRepo.ConsultarPontoDoDia(email)
	if err != nil {
		return "", err
	}

	pontos := make([]dominio.Ponto, len(registrosDoDia))
	for i, registro := range registrosDoDia {
		pontos[i] = *dominio.NewPonto(registro.Evento, registro.Timestamp)

		}
	}

	pontoDoDia := *dominio.NewPontoDoDia(email, pontos)*/

	c.pontoRepo.RegistrarPonto(email, ponto.Tipo, email)
	return "", nil
}