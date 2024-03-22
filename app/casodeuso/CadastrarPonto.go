package casodeuso

import "github.com/mvgv/lambda-registros/app/dominio"

type CadastrarPonto interface {
	CadastrarPontoDoDia(ponto dominio.Ponto, email string) (string, error)
}
