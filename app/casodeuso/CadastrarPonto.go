package casodeuso

type CadastrarPonto interface {
	CadastrarPontoDoDia(email, timestamp, evento string) (string, error)
}
