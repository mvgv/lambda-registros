package repositorio

import "github.com/mvgv/lambda-registros/app/infraestrutura/dto"

type PontoRepositorio interface {
	RegistrarPonto(email string, timestamp string, evento string) error
	ConsultarPontoDoDia(email string) (*dto.PontoDoDiaEntidade, error)
}
