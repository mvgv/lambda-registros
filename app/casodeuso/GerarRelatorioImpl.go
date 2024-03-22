package casodeuso

import "github.com/mvgv/lambda-registros/app/infraestrutura/mensagens"

type GerarRelatorioImpl struct {
	producer mensagens.Produtor
}

func NewGerarRelatorioImpl(producer mensagens.Produtor) *GerarRelatorioImpl {
	return &GerarRelatorioImpl{
		producer: producer,
	}
}

func (g *GerarRelatorioImpl) GerarRelatorioMensal(email, mes string) error {
	return nil
}
