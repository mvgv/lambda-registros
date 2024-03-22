package controlador

import (
	"github.com/mvgv/lambda-registros/app/apresentacao"
	"github.com/mvgv/lambda-registros/app/casodeuso"
)

type ConsultaPontoController struct {
	consultarClienteUC casodeuso.ConsultarCliente
	consultarPontoUC   casodeuso.ConsultarPonto
}

func NewConsultaPontoController(consultarClienteUC casodeuso.ConsultarCliente,
	consultarPontoUC casodeuso.ConsultarPonto) *ConsultaPontoController {
	return &ConsultaPontoController{
		consultarClienteUC: consultarClienteUC,
		consultarPontoUC:   consultarPontoUC,
	}
}

func (c *ConsultaPontoController) Handle(email string) (*apresentacao.PontoDoDiaRespostaRequisicao, error) {

	ponto, err := c.consultarPontoUC.ConsultarPontoDoDia(email)
	if err != nil {
		return nil, err
	}

	listaDePontos := make([]*apresentacao.PontoRespostaRequisicao, len(ponto.Registros))
	for i, ponto := range ponto.Registros {
		listaDePontos[i] = apresentacao.NewPontoRespostaRequisicao(ponto.Horario)
	}
	pontoResposta := apresentacao.NewPontoDoDiaRespostaRequisicao(listaDePontos)

	return pontoResposta, nil
}
