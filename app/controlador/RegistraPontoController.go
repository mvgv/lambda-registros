package controlador

import (
	"github.com/mvgv/lambda-registros/app/apresentacao"
	"github.com/mvgv/lambda-registros/app/casodeuso"
)

type RegistraPontoController struct {
	cadastrarPontoUC   casodeuso.CadastrarPonto
	consultarClienteUC casodeuso.ConsultarCliente
}

func NewRegistraPontoController(cadastrarPontoUC casodeuso.CadastrarPonto,
	consultarClienteUC casodeuso.ConsultarCliente) *RegistraPontoController {
	return &RegistraPontoController{
		cadastrarPontoUC:   cadastrarPontoUC,
		consultarClienteUC: consultarClienteUC,
	}
}

func (r *RegistraPontoController) Handle(apresentacao.PontoRequisicao) (string, error) {
	return "", nil
}
