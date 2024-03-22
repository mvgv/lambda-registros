package apresentacao

type PontoRequisicao struct {
	Email  string `json:"email"`
	Evento string `json:"evento"`
}

func NewPontoRequisicao(email, evento string) *PontoRequisicao {
	return &PontoRequisicao{
		Email:  email,
		Evento: evento,
	}
}
