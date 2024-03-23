package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mvgv/lambda-registros/app/apresentacao"
	"github.com/mvgv/lambda-registros/app/casodeuso"
	"github.com/mvgv/lambda-registros/app/controlador"
	"github.com/mvgv/lambda-registros/app/infraestrutura/mensagens"
	"github.com/mvgv/lambda-registros/app/infraestrutura/repositorio"
	"github.com/mvgv/lambda-registros/app/infraestrutura/servicos"
)

type Response struct {
	Message string `json:"message"`
}

func RegistrarPontoHandler(ctx context.Context, req events.APIGatewayProxyRequest,
	cadastrarPontoUC casodeuso.CadastrarPonto,
	consultaClienteUC casodeuso.ConsultarCliente) (events.APIGatewayProxyResponse, error) {

	var pontoDTO apresentacao.PontoRequisicao
	// TODO: Implementar a lógica de criação de cliente
	autorizador := apresentacao.NewValidaToken()
	permitido, _ := autorizador.AutorizarCliente(req.Headers["Authorization"], req.PathParameters["id_funcionario"])
	if !permitido {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "mensagem: Usuário não autorizado"}, nil
	}
	controller := controlador.NewRegistraPontoController(cadastrarPontoUC, consultaClienteUC)
	log.Printf("req.Body: %s\n", req.Body)

	err := json.Unmarshal([]byte(req.Body), &pontoDTO)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to unmarshal request: %v", err)
	}
	timestampRegistrado, evento, err := controller.Handle(pontoDTO.Email, time.Now().Format("2006-01-02T15:04:05"), pontoDTO.Evento)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to handle request: %v", err)
	}

	saidaPonto := apresentacao.NewPontoRespostaRequisicao(timestampRegistrado, evento)
	respBody, _ := json.Marshal(saidaPonto)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBody),
	}, nil
}

func RelatorioHandler(ctx context.Context, req events.APIGatewayProxyRequest,
	relatorioPontoUC casodeuso.GerarRelatorio,
	consultarClienteUC casodeuso.ConsultarCliente) (events.APIGatewayProxyResponse, error) {
	var pontoDto apresentacao.PontoRequisicao

	controller := controlador.NewGeraRelatorioController(relatorioPontoUC)
	err := json.Unmarshal([]byte(req.Body), &pontoDto)

	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to unmarshal request: %v", err)
	}

	autorizador := apresentacao.NewValidaToken()
	permitido, _ := autorizador.AutorizarCliente(req.Headers["Authorization"], pontoDto.Email)

	if !permitido {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "mensagem: Usuário não autorizado"}, nil
	}
	respBody, err := controller.Handle(pontoDto.Email)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "mensagem: Funcionario não encontrado"}, fmt.Errorf("failed to handle request: %v", err)
	}

	returnJson, _ := json.Marshal(apresentacao.NewMensagemRelatorio(respBody))

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnJson),
	}, nil

}

func ConsultaPontoHandler(ctx context.Context, req events.APIGatewayProxyRequest,
	consultarPontoUC casodeuso.ConsultarPonto,
	consultaClienteUC casodeuso.ConsultarCliente,
	calcularHorasTrabalhadasUC casodeuso.CalcularHorasTrabalhadas) (events.APIGatewayProxyResponse, error) {
	controller := controlador.NewConsultaPontoController(consultaClienteUC, consultarPontoUC, calcularHorasTrabalhadasUC)

	autorizador := apresentacao.NewValidaToken()
	permitido, _ := autorizador.AutorizarCliente(req.Headers["Authorization"], req.PathParameters["id_funcionario"])
	if !permitido {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusUnauthorized, Body: "mensagem: Usuário não autorizado"}, nil
	}

	respBody, err := controller.Handle(req.PathParameters["id_funcionario"])

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: "mensagem: Funcionario não encontrado"}, fmt.Errorf("failed to handle request: %v", err)
	}

	returnJson, _ := json.Marshal(respBody)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnJson),
	}, nil

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest,
	cadastrarPontoUC casodeuso.CadastrarPonto, consultarPontoUC casodeuso.ConsultarPonto,
	relatorioPontoUC casodeuso.GerarRelatorio, consultaClienteUC casodeuso.ConsultarCliente,
	caclularHorasTrabalhadasUC casodeuso.CalcularHorasTrabalhadas) (events.APIGatewayProxyResponse, error) {

	log.Printf("req.Path: %s\n", req.Path)
	switch req.HTTPMethod {
	case "POST":
		if strings.HasSuffix(req.Path, "/registros") {
			return RegistrarPontoHandler(ctx, req, cadastrarPontoUC, consultaClienteUC)
		} else {
			return RelatorioHandler(ctx, req, relatorioPontoUC, consultaClienteUC)
		}
	case "GET":
		return ConsultaPontoHandler(ctx, req, consultarPontoUC, consultaClienteUC, caclularHorasTrabalhadasUC)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       http.StatusText(http.StatusNotFound),
	}, nil
}

func main() {
	clienteService := servicos.NewServicoClienteImpl()             //falta implementar
	pontoRepository := repositorio.NewPontoRepositorioDynamoImpl() //falta implementar
	messageSNS := mensagens.NewProdutorSNSImpl()                   //falta implementar

	calculaHorasTrabalhadasUC := casodeuso.NewCalcularHorasTrabalhadasImpl()

	consultarClienteUC := casodeuso.NewConsultarClienteImpl(clienteService) //falta implementar
	cadastrarPontoUC := casodeuso.NewCadastrarPontoImpl(pontoRepository)    //falta implementar
	consultarPontoUC := casodeuso.NewConsultarPontoImpl(pontoRepository)    //falta implementar
	relatorioPontoUC := casodeuso.NewGerarRelatorioImpl(messageSNS)         //falta implementar

	lambda.Start(func(ctx context.Context, req map[string]interface{}) (interface{}, error) {
		fmt.Printf("req: %v\n", req)

		proxyRequestJSON, err := json.Marshal(req)
		if err != nil {
			return nil, fmt.Errorf("event type not supported")
		}
		var proxyRequestObj events.APIGatewayProxyRequest
		if err := json.Unmarshal(proxyRequestJSON, &proxyRequestObj); err != nil {
			return nil, fmt.Errorf("event type not supported")
		}
		fmt.Printf("proxyRequest: %v\n", proxyRequestObj)
		return Handler(ctx, proxyRequestObj, cadastrarPontoUC, consultarPontoUC, relatorioPontoUC, consultarClienteUC, calculaHorasTrabalhadasUC)

	})

}
