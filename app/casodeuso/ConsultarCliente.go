package casodeuso

type ConsultarCliente interface {
	ConsultarCliente(email string) (bool, error)
}
