package controlador

import (
	"time"

	"github.com/mvgv/lambda-registros/app/casodeuso"
)

type GeraRelatorio struct {
	gerarRelatorioUC casodeuso.GerarRelatorio
}

func NewGeraRelatorioController(gerarRelatorioUC casodeuso.GerarRelatorio) *GeraRelatorio {
	return &GeraRelatorio{
		gerarRelatorioUC: gerarRelatorioUC,
	}
}

func (g *GeraRelatorio) Handle(email string) (string, error) {
	lastMonth := time.Now().AddDate(0, -1, 0).Format("2006-01")
	err := g.gerarRelatorioUC.GerarRelatorioMensal(email, lastMonth)
	if err != nil {
		return "", err
	}
	return "Relatório gerado com sucesso consulte seu email", nil
}
