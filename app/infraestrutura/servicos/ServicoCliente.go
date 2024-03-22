package servicos

import (
	"github.com/mvgv/lambda-registros/app/infraestrutura/dto"
)

type ServicoCliente interface {
	GetCliente(email string) (*dto.ClienteRequisicao, error)
}
