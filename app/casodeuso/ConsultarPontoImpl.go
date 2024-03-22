package casodeuso

import (
	"github.com/mvgv/lambda-registros/app/dominio"
	"github.com/mvgv/lambda-registros/app/infraestrutura/repositorio"
)

type ConsultarPontoImpl struct {
	repositorio repositorio.PontoRepositorio
}

func NewConsultarPontoImpl(repositorio repositorio.PontoRepositorio) *ConsultarPontoImpl {
	return &ConsultarPontoImpl{
		repositorio: repositorio,
	}
}

func (c *ConsultarPontoImpl) ConsultarPontoDoDia(email string) (*dominio.PontoDoDia, error) {

	registrosDoDia, err := c.repositorio.ConsultarPontoDoDia(email)
	if err != nil {
		return nil, err
	}
	ponto := make([]dominio.Ponto, len(registrosDoDia))
	for i, registro := range registrosDoDia {
		ponto[i] = *dominio.NewPonto(string(registro[i][i]), string(registro[i][i]))
	}
	pontoDodia := dominio.NewPontoDoDia(email, ponto)

	return pontoDodia, nil

}
