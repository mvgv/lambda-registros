package controlador

import (
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

func (r *RegistraPontoController) Handle(email, timestamp, evento string) (string, string, error) {

	timestamp, evento, err := r.cadastrarPontoUC.CadastrarPontoDoDia(email, timestamp, evento)
	if err != nil {
		return "", "", err
	}
	return timestamp, evento, nil
}
