package casodeuso

import "github.com/mvgv/lambda-registros/app/dominio"

type ConsultarPonto interface {
	ConsultarPontoDoDia(email string) (*dominio.PontoDoDia, error)
}
