package casodeuso

import "github.com/mvgv/lambda-registros/app/infraestrutura/servicos"

type ConsultarClienteImpl struct {
	servicoCliente servicos.ServicoCliente
}

func NewConsultarClienteImpl(servicoCliente servicos.ServicoCliente) *ConsultarClienteImpl {
	return &ConsultarClienteImpl{
		servicoCliente: servicoCliente,
	}
}

func (c *ConsultarClienteImpl) ConsultarCliente(email string) (bool, error) {
	_, err := c.servicoCliente.GetCliente(email)
	if err != nil {
		return true, err
	}
	return false, nil
}
