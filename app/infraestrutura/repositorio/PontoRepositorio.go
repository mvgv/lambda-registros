package repositorio

type PontoRepositorio interface {
	RegistrarPonto(email string, timestamp string, evento string) error
	ConsultarPontoDoDia(email string) ([][]string, error)
}
