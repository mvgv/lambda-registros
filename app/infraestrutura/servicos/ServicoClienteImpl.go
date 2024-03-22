package servicos

import (
	"github.com/mvgv/lambda-registros/app/infraestrutura/dto"
)

type ServicoClienteImpl struct {
}

func NewServicoClienteImpl() *ServicoClienteImpl {
	return &ServicoClienteImpl{}
}

func (s *ServicoClienteImpl) GetCliente(email string) (*dto.ClienteRequisicao, error) {
	return dto.NewClienteRequisicao(email), nil
}
