package casodeuso

import (
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

func (c *CadastrarPontoImpl) CadastrarPontoDoDia(email, timestamp, evento string) (string, error) {

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

	err := c.pontoRepo.RegistrarPonto(email, timestamp, evento)
	if err != nil {
		return "", err
	}
	return timestamp, nil
}
