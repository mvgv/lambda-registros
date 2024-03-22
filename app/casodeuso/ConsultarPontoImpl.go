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
	return nil, nil
}
