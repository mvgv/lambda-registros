package apresentacao

type PontoRequisicao struct {
	Email     string `json:"email"`
	Timestamp string `json:"timestamp"`
}

func NewPontoRequisicao(email, timestamp string) *PontoRequisicao {
	return &PontoRequisicao{
		Email:     email,
		Timestamp: timestamp,
	}
}
