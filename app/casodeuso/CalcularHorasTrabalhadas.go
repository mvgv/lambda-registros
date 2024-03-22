package casodeuso

import "github.com/mvgv/lambda-registros/app/dominio"

type CalcularHorasTrabalhadas interface {
	CalcularHorasTrabalhadasNoDia(listaPontos []dominio.Ponto) (string, error)
}
