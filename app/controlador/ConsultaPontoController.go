package controlador

import (
	"fmt"

	"github.com/mvgv/lambda-registros/app/apresentacao"
	"github.com/mvgv/lambda-registros/app/casodeuso"
)

type ConsultaPontoController struct {
	consultarClienteUC         casodeuso.ConsultarCliente
	consultarPontoUC           casodeuso.ConsultarPonto
	calcularHorasTrabalhadasUC casodeuso.CalcularHorasTrabalhadas
}

func NewConsultaPontoController(consultarClienteUC casodeuso.ConsultarCliente,
	consultarPontoUC casodeuso.ConsultarPonto, calcularHorasTrabalhadasUC casodeuso.CalcularHorasTrabalhadas) *ConsultaPontoController {
	return &ConsultaPontoController{
		consultarClienteUC:         consultarClienteUC,
		consultarPontoUC:           consultarPontoUC,
		calcularHorasTrabalhadasUC: calcularHorasTrabalhadasUC,
	}
}

func (c *ConsultaPontoController) Handle(email string) (*apresentacao.PontoDoDiaRespostaRequisicao, error) {
	fmt.Println("ConsultaPontoController.Handle()")
	ponto, err := c.consultarPontoUC.ConsultarPontoDoDia(email)
	if err != nil {
		return nil, err
	}

	horasTrabalhas, _ := c.calcularHorasTrabalhadasUC.CalcularHorasTrabalhadasNoDia(ponto.Registros)

	listaDePontos := make([]*apresentacao.PontoRespostaRequisicao, len(ponto.Registros))
	for i, ponto := range ponto.Registros {
		listaDePontos[i] = apresentacao.NewPontoRespostaRequisicao(ponto.Horario, ponto.Tipo)
	}
	pontoResposta := apresentacao.NewPontoDoDiaRespostaRequisicao(listaDePontos, horasTrabalhas)

	return pontoResposta, nil
}
