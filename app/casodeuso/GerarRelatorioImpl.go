package casodeuso

import (
	"github.com/mvgv/lambda-registros/app/infraestrutura/dto"
	"github.com/mvgv/lambda-registros/app/infraestrutura/mensagens"
)

type GerarRelatorioImpl struct {
	producer mensagens.Produtor
}

func NewGerarRelatorioImpl(producer mensagens.Produtor) *GerarRelatorioImpl {
	return &GerarRelatorioImpl{
		producer: producer,
	}
}

func (g *GerarRelatorioImpl) GerarRelatorioMensal(email, mes string) error {
	mensagem := dto.NewSolicitacaoRelatorio(email, mes)
	err := g.producer.EnviarMensagem(mensagem)
	if err != nil {
		return err
	}
	return nil
}
